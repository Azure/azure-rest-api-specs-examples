/** Samples for BackupInstances ResumeProtection. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/ResumeProtection.json
     */
    /**
     * Sample code: ResumeProtection.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void resumeProtection(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().resumeProtection("testrg", "testvault", "testbi", com.azure.core.util.Context.NONE);
    }
}
