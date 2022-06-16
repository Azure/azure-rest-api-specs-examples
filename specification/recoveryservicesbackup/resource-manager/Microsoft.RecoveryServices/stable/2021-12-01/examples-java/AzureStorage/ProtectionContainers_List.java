import com.azure.core.util.Context;

/** Samples for BackupProtectionContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/AzureStorage/ProtectionContainers_List.json
     */
    /**
     * Sample code: List Backup Protection Containers.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listBackupProtectionContainers(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupProtectionContainers()
            .list("testVault", "testRg", "backupManagementType eq 'AzureWorkload'", Context.NONE);
    }
}
