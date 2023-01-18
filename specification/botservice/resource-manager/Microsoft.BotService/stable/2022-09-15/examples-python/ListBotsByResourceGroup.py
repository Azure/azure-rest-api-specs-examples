from azure.identity import DefaultAzureCredential
from azure.mgmt.botservice import AzureBotService

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-botservice
# USAGE
    python list_bots_by_resource_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureBotService(
        credential=DefaultAzureCredential(),
        subscription_id="subscription-id",
    )

    response = client.bots.list_by_resource_group(
        resource_group_name="OneResourceGroupName",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/ListBotsByResourceGroup.json
if __name__ == "__main__":
    main()
