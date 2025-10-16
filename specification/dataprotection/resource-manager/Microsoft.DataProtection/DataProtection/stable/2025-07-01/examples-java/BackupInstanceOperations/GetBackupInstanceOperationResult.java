
/**
 * Samples for BackupInstances GetBackupInstanceOperationResult.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BackupInstanceOperations/GetBackupInstanceOperationResult.json
     */
    /**
     * Sample code: Get BackupInstanceOperationResult.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        getBackupInstanceOperationResult(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().getBackupInstanceOperationResultWithResponse("SampleResourceGroup", "swaggerExample",
            "testInstance1",
            "YWUzNDFkMzQtZmM5OS00MmUyLWEzNDMtZGJkMDIxZjlmZjgzOzdmYzBiMzhmLTc2NmItNDM5NS05OWQ1LTVmOGEzNzg4MWQzNA==",
            com.azure.core.util.Context.NONE);
    }
}
