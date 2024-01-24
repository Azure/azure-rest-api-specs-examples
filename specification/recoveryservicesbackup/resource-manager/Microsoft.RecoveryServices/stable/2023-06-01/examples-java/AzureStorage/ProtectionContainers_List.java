
/**
 * Samples for BackupProtectionContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/
     * AzureStorage/ProtectionContainers_List.json
     */
    /**
     * Sample code: List Backup Protection Containers.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listBackupProtectionContainers(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupProtectionContainers().list("testVault", "testRg", "backupManagementType eq 'AzureWorkload'",
            com.azure.core.util.Context.NONE);
    }
}
