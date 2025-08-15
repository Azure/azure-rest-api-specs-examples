
/**
 * Samples for BackupVaultOperationResults Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2025-07-01/examples/VaultCRUD/
     * GetOperationResultPatch.json
     */
    /**
     * Sample code: GetOperationResult Patch.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getOperationResultPatch(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupVaultOperationResults().getWithResponse("SampleResourceGroup", "swaggerExample",
            "YWUzNDFkMzQtZmM5OS00MmUyLWEzNDMtZGJkMDIxZjlmZjgzOzdmYzBiMzhmLTc2NmItNDM5NS05OWQ1LTVmOGEzNzg4MWQzNA==",
            com.azure.core.util.Context.NONE);
    }
}
