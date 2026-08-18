
import com.azure.resourcemanager.recoveryservicesbackup.models.AccessType;
import com.azure.resourcemanager.recoveryservicesbackup.models.AcquireStorageAccountLock;
import com.azure.resourcemanager.recoveryservicesbackup.models.AzureStorageContainer;
import com.azure.resourcemanager.recoveryservicesbackup.models.BackupManagementType;
import com.azure.resourcemanager.recoveryservicesbackup.models.IdentityInfo;
import com.azure.resourcemanager.recoveryservicesbackup.models.OperationType;

/**
 * Samples for ProtectionContainers Register.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/AzureStorage/ProtectionContainers_ReRegister_SwitchToUAMI.json
     */
    /**
     * Sample code: Re-register Azure Storage ProtectionContainers switching to User Assigned Managed Identity.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void reRegisterAzureStorageProtectionContainersSwitchingToUserAssignedManagedIdentity(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionContainers().define("StorageContainer;Storage;SwaggerTestRg;swaggertestsa")
            .withExistingBackupFabric("swaggertestvault", "SwaggerTestRg", "Azure")
            .withProperties(new AzureStorageContainer().withFriendlyName("swaggertestsa")
                .withBackupManagementType(BackupManagementType.AZURE_STORAGE)
                .withSourceResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/SwaggerTestRg/providers/Microsoft.Storage/storageAccounts/swaggertestsa")
                .withAcquireStorageAccountLock(AcquireStorageAccountLock.ACQUIRE)
                .withOperationType(OperationType.REREGISTER).withAccessType(AccessType.IDENTITY_BASED)
                .withIdentityInfo(new IdentityInfo().withIsSystemAssignedIdentity(false).withManagedIdentityResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/SwaggerTestRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/swaggertestuami")))
            .create();
    }
}
