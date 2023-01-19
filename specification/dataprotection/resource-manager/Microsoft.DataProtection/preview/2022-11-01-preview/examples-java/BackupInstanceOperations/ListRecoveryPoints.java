/** Samples for RecoveryPoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/ListRecoveryPoints.json
     */
    /**
     * Sample code: List Recovery Points in a Vault.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void listRecoveryPointsInAVault(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .recoveryPoints()
            .list(
                "000pikumar",
                "PratikPrivatePreviewVault1",
                "testInstance1",
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
