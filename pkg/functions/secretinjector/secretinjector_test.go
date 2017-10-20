///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package secretinjector

import (
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/functions/mocks"

	secretclient "gitlab.eng.vmware.com/serverless/serverless/pkg/secret-store/gen/client"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/secret-store/gen/client/secret"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/secret-store/gen/models"
)

func TestImpl_GetMiddleware(t *testing.T) {

	expectedSecretName := "testSecret"
	expectedSecretValue := models.SecretValue{"secret1": "value1", "secret2": "value2"}
	expectedSecrets := map[string]interface{}{
		"testSecret": expectedSecretValue,
	}

	mockedTransport := &mocks.ClientTransport{}
	mockedTransport.On("Submit", mock.Anything).Return(
		&secret.GetSecretOK{
			Payload: &models.Secret{
				Name:    &expectedSecretName,
				Secrets: expectedSecretValue,
			}}, nil)

	secretStore := secretclient.New(mockedTransport, strfmt.Default)

	injector := New(secretStore)

	cookie := "testCookie"
	input := make(map[string]interface{})
	input["arg1"] = "hello"
	input["arg2"] = "world"
	echoWithSecrets := func(input map[string]interface{}) (map[string]interface{}, error) {
		output := input
		output["secrets"] = input["_meta"].(map[string]interface{})["secrets"]
		return output, nil
	}
	output, err := injector.GetMiddleware([]string{expectedSecretName}, cookie)(echoWithSecrets)(input)
	assert.NoError(t, err)
	assert.Nil(t, output["_meta"])
	assert.Equal(t, expectedSecrets, output["secrets"])
	assert.Equal(t, input["arg1"], output["arg1"])
	assert.Equal(t, input["arg2"], output["arg2"])
}
