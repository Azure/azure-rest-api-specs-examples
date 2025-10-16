
/**
 * Samples for RecoveryPoints List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BackupInstanceOperations/ListRecoveryPoints.json
     */
    /**
     * Sample code: List Recovery Points in a Vault.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        listRecoveryPointsInAVault(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.recoveryPoints().list("000pikumar", "PratikPrivatePreviewVault1", "testInstance1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
