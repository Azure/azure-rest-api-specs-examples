from azure.identity import DefaultAzureCredential
from azure.mgmt.healthbot import HealthBotMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-healthbot
# USAGE
    python bot_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HealthBotMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.bots.begin_create(
        resource_group_name="healthbotClient",
        bot_name="samplebotname",
        parameters={
            "identity": {
                "type": "SystemAssigned, UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/subscription-id/resourcegroups/myrg/providers/microsoft.managedidentity/userassignedidentities/my-mi": {},
                    "/subscriptions/subscription-id/resourcegroups/myrg/providers/microsoft.managedidentity/userassignedidentities/my-mi2": {},
                },
            },
            "location": "East US",
            "sku": {"name": "F0"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2022-08-08/examples/ResourceCreationPut.json
if __name__ == "__main__":
    main()
