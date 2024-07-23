
/**
 * Samples for BackupInstances List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * BackupInstanceOperations/ListBackupInstances.json
     */
    /**
     * Sample code: List BackupInstances in a Vault.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        listBackupInstancesInAVault(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().list("000pikumar", "PratikPrivatePreviewVault1", com.azure.core.util.Context.NONE);
    }
}
