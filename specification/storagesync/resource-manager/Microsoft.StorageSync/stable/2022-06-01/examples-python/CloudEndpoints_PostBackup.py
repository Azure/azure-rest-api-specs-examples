from azure.identity import DefaultAzureCredential
from azure.mgmt.storagesync import MicrosoftStorageSync

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storagesync
# USAGE
    python cloud_endpoints_post_backup.py

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

    response = client.cloud_endpoints.begin_post_backup(
        resource_group_name="SampleResourceGroup_1",
        storage_sync_service_name="SampleStorageSyncService_1",
        sync_group_name="SampleSyncGroup_1",
        cloud_endpoint_name="SampleCloudEndpoint_1",
        parameters={"azureFileShare": "https://sampleserver.file.core.test-cint.azure-test.net/sampleFileShare"},
    ).result()
    print(response)


# x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/CloudEndpoints_PostBackup.json
if __name__ == "__main__":
    main()
