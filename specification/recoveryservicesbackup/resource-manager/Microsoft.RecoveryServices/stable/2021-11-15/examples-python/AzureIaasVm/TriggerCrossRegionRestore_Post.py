from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicesbackup.passivestamp import RecoveryServicesBackupPassiveClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesbackup
# USAGE
    python trigger_cross_region_restore_post.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RecoveryServicesBackupPassiveClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    client.cross_region_restore.begin_trigger(
        azure_region="southeastasia",
        parameters={
            "crossRegionRestoreAccessDetails": {
                "accessTokenString": "**********",
                "backupManagementType": "AzureIaasVM",
                "containerName": "iaasvmcontainerv2;srinivasccyrg;sriniccylinux",
                "containerType": "IaasVMContainer",
                "coordinatorServiceStampUri": "https://pod01-coord1.ccy.backup.windowsazure.com",
                "datasourceContainerName": "iaasvmcontainerv2;srinivasccyrg;sriniccylinux",
                "datasourceId": "1142937031",
                "datasourceName": "sriniccylinux",
                "datasourceType": "VM",
                "objectType": "WorkloadCrrAccessToken",
                "protectionServiceStampId": "90d98224-2ac6-4bda-9f35-33fb22841f2a",
                "protectionServiceStampUri": "https://pod01-prot1-int.ccy.backup.windowsazure.com",
                "recoveryPointId": "87178355392716",
                "recoveryPointTime": "10/9/2019 6:05:54 PM",
                "resourceGroupName": "srinivasccyrg",
                "resourceId": "1330837906418138160",
                "resourceName": "sriniccyvault",
                "subscriptionId": "f2edfd5d-5496-4683-b94f-b3588c579009",
                "tokenExtendedInformation": '<IaaSVMRecoveryPointMetadataBase xmlns:i="http://www.w3.org/2001/XMLSchema-instance" i:type="IaaSVMRecoveryPointMetadata_V2015_09" xmlns="http://windowscloudbackup.com/CloudCommon/V2011_09"><MetadataVersion>V2015_09</MetadataVersion><ContainerType i:nil="true" /><InstantRpGCId>f2edfd5d-5496-4683-b94f-b3588c579009;AzureBackup_sriniccylinux_1142937031;AzureBackup_20191009_060554;AzureBackupRG_centraluseuap_1</InstantRpGCId><IsBlockBlobEnabled>true</IsBlockBlobEnabled><IsManagedVirtualMachine>true</IsManagedVirtualMachine><OriginalSAOption>false</OriginalSAOption><OsType>Linux</OsType><ReadMetadaFromConfigBlob i:nil="true" /><RecoveryPointConsistencyType>FileSystemConsistent</RecoveryPointConsistencyType><RpDiskDetails i:nil="true" /><SourceIaaSVMRPKeyAndSecret i:nil="true" /><SourceIaaSVMStorageType>PremiumVMOnPremiumStorage</SourceIaaSVMStorageType><VMSizeDescription>Standard_D2s_v3</VMSizeDescription></IaaSVMRecoveryPointMetadataBase>',
            },
            "restoreRequest": {
                "affinityGroup": "",
                "createNewCloudService": False,
                "encryptionDetails": {"encryptionEnabled": False},
                "identityInfo": {
                    "isSystemAssignedIdentity": False,
                    "managedIdentityResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/asmaskarRG1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/asmaskartestmsi",
                },
                "objectType": "IaasVMRestoreRequest",
                "originalStorageAccountOption": False,
                "recoveryPointId": "87178355392716",
                "recoveryType": "AlternateLocation",
                "region": "eastus2euap",
                "sourceResourceId": "/subscriptions/f2edfd5d-5496-4683-b94f-b3588c579009/resourceGroups/srinivasccyrg/providers/Microsoft.Compute/virtualMachines/sriniccylinux",
                "storageAccountId": "/subscriptions/f2edfd5d-5496-4683-b94f-b3588c579009/resourceGroups/00prjaiTestRg1/providers/Microsoft.Storage/storageAccounts/00prjaitestrg1disks993",
                "subnetId": "/subscriptions/f2edfd5d-5496-4683-b94f-b3588c579009/resourceGroups/00networkAcklVaultCCY/providers/Microsoft.Network/virtualNetworks/00networkAcklVaultCCY-vnet/subnets/default",
                "targetDomainNameId": None,
                "targetResourceGroupId": "/subscriptions/f2edfd5d-5496-4683-b94f-b3588c579009/resourceGroups/00networkAckl",
                "targetVirtualMachineId": "/subscriptions/f2edfd5d-5496-4683-b94f-b3588c579009/resourceGroups/00networkAckl/providers/Microsoft.Compute/virtualMachines/gaallaVM",
                "virtualNetworkId": "/subscriptions/f2edfd5d-5496-4683-b94f-b3588c579009/resourceGroups/00networkAcklVaultCCY/providers/Microsoft.Network/virtualNetworks/00networkAcklVaultCCY-vnet",
                "zones": ["2"],
            },
        },
    ).result()


# x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-11-15/examples/AzureIaasVm/TriggerCrossRegionRestore_Post.json
if __name__ == "__main__":
    main()
