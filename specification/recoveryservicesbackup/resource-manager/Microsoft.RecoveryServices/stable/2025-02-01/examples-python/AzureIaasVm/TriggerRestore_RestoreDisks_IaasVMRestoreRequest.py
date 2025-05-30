from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicesbackup.activestamp import RecoveryServicesBackupClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesbackup
# USAGE
    python trigger_restore_restore_disks_iaas_vm_restore_request.py

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

    client.restores.begin_trigger(
        vault_name="testVault",
        resource_group_name="netsdktestrg",
        fabric_name="Azure",
        container_name="IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
        protected_item_name="VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
        recovery_point_id="348916168024334",
        parameters={
            "properties": {
                "createNewCloudService": True,
                "encryptionDetails": {"encryptionEnabled": False},
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
                "storageAccountId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testingRg/providers/Microsoft.Storage/storageAccounts/testAccount",
                "targetDiskNetworkAccessSettings": {
                    "targetDiskAccessId": "/subscriptions/e7a191f5-713c-4bdb-b5e4-cf3dd90230ef/resourceGroups/arpja/providers/Microsoft.Compute/diskAccesses/arpja-diskaccess-ccy",
                    "targetDiskNetworkAccessOption": "EnablePrivateAccessForAllDisks",
                },
            }
        },
    ).result()


# x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureIaasVm/TriggerRestore_RestoreDisks_IaasVMRestoreRequest.json
if __name__ == "__main__":
    main()
