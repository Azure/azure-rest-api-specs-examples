from azure.identity import DefaultAzureCredential
from azure.mgmt.storagesync import MicrosoftStorageSync

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storagesync
# USAGE
    python private_link_resources_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MicrosoftStorageSync(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.private_link_resources.list_by_storage_sync_service(
        resource_group_name="res6977",
        storage_sync_service_name="sss2527",
    )
    print(response)


# x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/PrivateLinkResources_List.json
if __name__ == "__main__":
    main()
