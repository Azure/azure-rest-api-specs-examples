from azure.identity import DefaultAzureCredential

from azure.mgmt.storageactions import StorageActionsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storageactions
# USAGE
    python perform_storage_task_actions_preview.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageActionsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.storage_tasks.preview_actions(
        location="eastus",
        parameters={
            "properties": {
                "action": {"elseBlockExists": True, "if": {"condition": "[[equals(AccessTier, 'Hot')]]"}},
                "blobs": [
                    {
                        "metadata": [{"key": "mKey1", "value": "mValue1"}],
                        "name": "folder1/file1.txt",
                        "properties": [
                            {"key": "Creation-Time", "value": "Wed, 07 Jun 2023 05:23:29 GMT"},
                            {"key": "Last-Modified", "value": "Wed, 07 Jun 2023 05:23:29 GMT"},
                            {"key": "Etag", "value": "0x8DB67175454D36D"},
                            {"key": "Content-Length", "value": "38619"},
                            {"key": "Content-Type", "value": "text/xml"},
                            {"key": "Content-Encoding", "value": ""},
                            {"key": "Content-Language", "value": ""},
                            {"key": "Content-CRC64", "value": ""},
                            {"key": "Content-MD5", "value": "njr6iDrmU9+FC89WMK22EA=="},
                            {"key": "Cache-Control", "value": ""},
                            {"key": "Content-Disposition", "value": ""},
                            {"key": "BlobType", "value": "BlockBlob"},
                            {"key": "AccessTier", "value": "Hot"},
                            {"key": "AccessTierInferred", "value": "true"},
                            {"key": "LeaseStatus", "value": "unlocked"},
                            {"key": "LeaseState", "value": "available"},
                            {"key": "ServerEncrypted", "value": "true"},
                            {"key": "TagCount", "value": "1"},
                        ],
                        "tags": [{"key": "tKey1", "value": "tValue1"}],
                    },
                    {
                        "metadata": [{"key": "mKey2", "value": "mValue2"}],
                        "name": "folder2/file1.txt",
                        "properties": [
                            {"key": "Creation-Time", "value": "Wed, 06 Jun 2023 05:23:29 GMT"},
                            {"key": "Last-Modified", "value": "Wed, 06 Jun 2023 05:23:29 GMT"},
                            {"key": "Etag", "value": "0x6FB67175454D36D"},
                        ],
                        "tags": [{"key": "tKey2", "value": "tValue2"}],
                    },
                ],
                "container": {
                    "metadata": [{"key": "mContainerKey1", "value": "mContainerValue1"}],
                    "name": "firstContainer",
                },
            }
        },
    )
    print(response)


# x-ms-original-file: 2023-01-01/misc/PerformStorageTaskActionsPreview.json
if __name__ == "__main__":
    main()
