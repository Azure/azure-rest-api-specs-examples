from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.compute import ComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-compute
# USAGE
    python disk_create_by_importing_blob_from_adifferent_subscription.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ComputeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.disks.begin_create_or_update(
        resource_group_name="myResourceGroup",
        disk_name="myDisk",
        disk={
            "location": "West US",
            "properties": {
                "creationData": {
                    "createOption": "Import",
                    "sourceUri": "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
                    "storageAccountId": "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount",
                }
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskExamples/Disk_Create_ByImportingBlobFromADifferentSubscription.json
if __name__ == "__main__":
    main()
