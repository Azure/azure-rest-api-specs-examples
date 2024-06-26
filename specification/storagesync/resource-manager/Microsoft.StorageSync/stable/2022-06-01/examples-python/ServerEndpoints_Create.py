from azure.identity import DefaultAzureCredential
from azure.mgmt.storagesync import MicrosoftStorageSync

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storagesync
# USAGE
    python server_endpoints_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MicrosoftStorageSync(
        credential=DefaultAzureCredential(),
        subscription_id="52b8da2f-61e0-4a1f-8dde-336911f367fb",
    )

    response = client.server_endpoints.begin_create(
        resource_group_name="SampleResourceGroup_1",
        storage_sync_service_name="SampleStorageSyncService_1",
        sync_group_name="SampleSyncGroup_1",
        server_endpoint_name="SampleServerEndpoint_1",
        parameters={
            "properties": {
                "cloudTiering": "off",
                "initialDownloadPolicy": "NamespaceThenModifiedFiles",
                "initialUploadPolicy": "ServerAuthoritative",
                "localCacheMode": "UpdateLocallyCachedFiles",
                "offlineDataTransfer": "on",
                "offlineDataTransferShareName": "myfileshare",
                "serverLocalPath": "D:\\SampleServerEndpoint_1",
                "serverResourceId": "/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/080d4133-bdb5-40a0-96a0-71a6057bfe9a",
                "tierFilesOlderThanDays": 0,
                "volumeFreeSpacePercent": 100,
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/ServerEndpoints_Create.json
if __name__ == "__main__":
    main()
