/** Samples for DeletedBackupInstances Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/DeletedBackupInstanceOperations/GetDeletedBackupInstance.json
     */
    /**
     * Sample code: Get DeletedBackupInstance.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getDeletedBackupInstance(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .deletedBackupInstances()
            .getWithResponse(
                "000pikumar", "PratikPrivatePreviewVault1", "testInstance1", com.azure.core.util.Context.NONE);
    }
}
