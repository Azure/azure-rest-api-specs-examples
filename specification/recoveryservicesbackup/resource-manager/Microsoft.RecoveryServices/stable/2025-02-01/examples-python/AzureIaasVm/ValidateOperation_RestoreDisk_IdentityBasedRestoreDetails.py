from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicesbackup.activestamp import RecoveryServicesBackupClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesbackup
# USAGE
    python validate_operation_restore_disk_identity_based_restore_details.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RecoveryServicesBackupClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.operation.validate(
        vault_name="testVault",
        resource_group_name="testRG",
        parameters={
            "id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testVault/providers/Microsoft.RecoveryServices/vaults/testVault/backupFabrics/Azure/protectionContainers/IaasVMContainer;iaasvmcontainerv2;testRG;testvmName/protectedItems/VM;iaasvmcontainerv2;testRG;testvmName/recoveryPoints/348916168024334",
            "properties": {
                "objectType": "ValidateIaasVMRestoreOperationRequest",
                "restoreRequest": {
                    "createNewCloudService": True,
                    "encryptionDetails": {"encryptionEnabled": False},
                    "identityBasedRestoreDetails": {
                        "targetStorageAccountId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testingRg/providers/Microsoft.Storage/storageAccounts/testAccount"
                    },
                    "identityInfo": {
                        "isSystemAssignedIdentity": False,
                        "managedIdentityResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/asmaskarRG1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/asmaskartestmsi",
                    },
                    "objectType": "IaasVMRestoreRequest",
                    "originalStorageAccountOption": False,
                    "recoveryPointId": "348916168024334",
                    "recoveryType": "RestoreDisks",
                    "region": "southeastasia",
                    "sourceResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1",
                },
            },
        },
    )
    print(response)


# x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureIaasVm/ValidateOperation_RestoreDisk_IdentityBasedRestoreDetails.json
if __name__ == "__main__":
    main()
