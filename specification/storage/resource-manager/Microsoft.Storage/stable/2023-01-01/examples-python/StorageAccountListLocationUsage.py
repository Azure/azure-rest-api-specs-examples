from azure.identity import DefaultAzureCredential
from azure.mgmt.storage import StorageManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storage
# USAGE
    python storage_account_list_location_usage.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.usages.list_by_location(
        location="eastus2(stage)",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountListLocationUsage.json
if __name__ == "__main__":
    main()
