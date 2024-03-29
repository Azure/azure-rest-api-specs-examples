from azure.identity import DefaultAzureCredential
from azure.mgmt.machinelearningservices import MachineLearningServicesMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-machinelearningservices
# USAGE
    python create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MachineLearningServicesMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.jobs.create_or_update(
        resource_group_name="test-rg",
        workspace_name="my-aml-workspace",
        id="string",
        body={
            "properties": {
                "codeId": "string",
                "command": "string",
                "computeId": "string",
                "description": "string",
                "displayName": "string",
                "distribution": {"distributionType": "TensorFlow", "parameterServerCount": 1, "workerCount": 1},
                "environmentId": "string",
                "environmentVariables": {"string": "string"},
                "experimentName": "string",
                "identity": {"identityType": "AMLToken"},
                "inputs": {"string": {"description": "string", "jobInputType": "literal", "value": "string"}},
                "jobType": "Command",
                "limits": {"jobLimitsType": "Command", "timeout": "PT5M"},
                "outputs": {
                    "string": {
                        "description": "string",
                        "jobOutputType": "uri_file",
                        "mode": "ReadWriteMount",
                        "uri": "string",
                    }
                },
                "properties": {"string": "string"},
                "resources": {
                    "instanceCount": 1,
                    "instanceType": "string",
                    "properties": {"string": {"e6b6493e-7d5e-4db3-be1e-306ec641327e": None}},
                },
                "services": {
                    "string": {
                        "endpoint": "string",
                        "jobServiceType": "string",
                        "port": 1,
                        "properties": {"string": "string"},
                    }
                },
                "tags": {"string": "string"},
            }
        },
    )
    print(response)


# x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2023-04-01/examples/Job/CommandJob/createOrUpdate.json
if __name__ == "__main__":
    main()
