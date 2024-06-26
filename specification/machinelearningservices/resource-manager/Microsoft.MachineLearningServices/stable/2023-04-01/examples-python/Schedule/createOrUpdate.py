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

    response = client.schedules.begin_create_or_update(
        resource_group_name="test-rg",
        workspace_name="my-aml-workspace",
        name="string",
        body={
            "properties": {
                "action": {
                    "actionType": "InvokeBatchEndpoint",
                    "endpointInvocationDefinition": {"9965593e-526f-4b89-bb36-761138cf2794": None},
                },
                "description": "string",
                "displayName": "string",
                "isEnabled": False,
                "properties": {"string": "string"},
                "tags": {"string": "string"},
                "trigger": {
                    "endTime": "string",
                    "expression": "string",
                    "startTime": "string",
                    "timeZone": "string",
                    "triggerType": "Cron",
                },
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2023-04-01/examples/Schedule/createOrUpdate.json
if __name__ == "__main__":
    main()
