/** Samples for DeletedBackupInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/DeletedBackupInstanceOperations/ListDeletedBackupInstances.json
     */
    /**
     * Sample code: List DeletedBackupInstances in a Vault.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void listDeletedBackupInstancesInAVault(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .deletedBackupInstances()
            .list("000pikumar", "PratikPrivatePreviewVault1", com.azure.core.util.Context.NONE);
    }
}
