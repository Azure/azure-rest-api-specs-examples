from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.compute import ComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-compute
# USAGE
    python disk_encryption_set_update.py

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

    response = client.disk_encryption_sets.begin_update(
        resource_group_name="myResourceGroup",
        disk_encryption_set_name="myDiskEncryptionSet",
        disk_encryption_set={
            "properties": {
                "activeKey": {
                    "keyUrl": "https://myvmvault.vault-int.azure-int.net/keys/keyName/keyVersion",
                    "sourceVault": {
                        "id": "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"
                    },
                },
                "encryptionType": "EncryptionAtRestWithCustomerKey",
            },
            "tags": {"department": "Development", "project": "Encryption"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Update.json
if __name__ == "__main__":
    main()
