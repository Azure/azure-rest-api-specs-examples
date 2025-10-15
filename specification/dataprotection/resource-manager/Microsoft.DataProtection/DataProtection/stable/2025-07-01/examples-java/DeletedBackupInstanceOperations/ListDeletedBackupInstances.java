
/**
 * Samples for DeletedBackupInstances List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DeletedBackupInstanceOperations/ListDeletedBackupInstances.json
     */
    /**
     * Sample code: List DeletedBackupInstances in a Vault.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        listDeletedBackupInstancesInAVault(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.deletedBackupInstances().list("000pikumar", "PratikPrivatePreviewVault1",
            com.azure.core.util.Context.NONE);
    }
}
