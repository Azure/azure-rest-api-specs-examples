from azure.identity import DefaultAzureCredential
from azure.mgmt.botservice import AzureBotService

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-botservice
# USAGE
    python list_connections_by_bot_service.py

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

    response = client.bot_connection.list_by_bot_service(
        resource_group_name="OneResourceGroupName",
        resource_name="samplebotname",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/ListConnectionsByBotService.json
if __name__ == "__main__":
    main()
