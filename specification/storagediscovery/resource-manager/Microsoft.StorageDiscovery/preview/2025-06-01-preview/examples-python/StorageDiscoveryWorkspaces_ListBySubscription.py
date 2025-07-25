from azure.identity import DefaultAzureCredential

from azure.mgmt.storagediscovery import StorageDiscoveryMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storagediscovery
# USAGE
    python storage_discovery_workspaces_list_by_subscription.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageDiscoveryMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.storage_discovery_workspaces.list_by_subscription()
    for item in response:
        print(item)


# x-ms-original-file: 2025-06-01-preview/StorageDiscoveryWorkspaces_ListBySubscription.json
if __name__ == "__main__":
    main()
