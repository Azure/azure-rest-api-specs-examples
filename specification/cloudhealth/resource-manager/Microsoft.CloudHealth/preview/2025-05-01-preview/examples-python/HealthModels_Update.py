from azure.identity import DefaultAzureCredential

from azure.mgmt.cloudhealth import CloudHealthMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cloudhealth
# USAGE
    python health_models_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CloudHealthMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.health_models.begin_update(
        resource_group_name="rgopenapi",
        health_model_name="model1",
        properties={
            "identity": {
                "type": "SystemAssigned, UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/4980D7D5-4E07-47AD-AD34-E76C6BC9F061/resourceGroups/rgopenapi/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ua1": {}
                },
            },
            "tags": {"key21": "menfkmseplchh"},
        },
    ).result()
    print(response)


# x-ms-original-file: 2025-05-01-preview/HealthModels_Update.json
if __name__ == "__main__":
    main()
