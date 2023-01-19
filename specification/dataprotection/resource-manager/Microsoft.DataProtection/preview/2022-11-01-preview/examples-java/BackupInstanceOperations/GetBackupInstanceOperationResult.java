/** Samples for BackupInstances GetBackupInstanceOperationResult. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/BackupInstanceOperations/GetBackupInstanceOperationResult.json
     */
    /**
     * Sample code: Get BackupInstanceOperationResult.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getBackupInstanceOperationResult(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .backupInstances()
            .getBackupInstanceOperationResultWithResponse(
                "SampleResourceGroup",
                "swaggerExample",
                "testInstance1",
                "YWUzNDFkMzQtZmM5OS00MmUyLWEzNDMtZGJkMDIxZjlmZjgzOzdmYzBiMzhmLTc2NmItNDM5NS05OWQ1LTVmOGEzNzg4MWQzNA==",
                com.azure.core.util.Context.NONE);
    }
}
