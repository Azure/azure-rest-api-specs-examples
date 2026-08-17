
/**
 * Samples for BackupProtectionContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/AzureStorage/ProtectionContainers_List_WithAccessType.json
     */
    /**
     * Sample code: List Backup Protection Containers with Access Type.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listBackupProtectionContainersWithAccessType(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupProtectionContainers().list("swaggertestvault", "SwaggerTestRg",
            "backupManagementType eq 'AzureStorage'", com.azure.core.util.Context.NONE);
    }
}
