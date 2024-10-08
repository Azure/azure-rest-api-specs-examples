from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.compute import ComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-compute
# USAGE
    python disk_encryption_set_create_with_key_vault_from_adifferent_tenant.py

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

    response = client.disk_encryption_sets.begin_create_or_update(
        resource_group_name="myResourceGroup",
        disk_encryption_set_name="myDiskEncryptionSet",
        disk_encryption_set={
            "identity": {
                "type": "UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}": {}
                },
            },
            "location": "West US",
            "properties": {
                "activeKey": {"keyUrl": "https://myvaultdifferenttenant.vault-int.azure-int.net/keys/{key}"},
                "encryptionType": "EncryptionAtRestWithCustomerKey",
                "federatedClientId": "00000000-0000-0000-0000-000000000000",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Create_WithKeyVaultFromADifferentTenant.json
if __name__ == "__main__":
    main()
