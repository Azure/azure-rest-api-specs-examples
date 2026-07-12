
/**
 * Samples for ResourceProvider GetOperationResult.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/GetOperationResult.json
     */
    /**
     * Sample code: Get Operation Result.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void getOperationResult(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.resourceProviders().getOperationResultWithResponse("HelloWorld", "swaggerExample",
            "YWUzNDFkMzQtZmM5OS00MmUyLWEzNDMtZGJkMDIxZjlmZjgzOzdmYzBiMzhmLTc2NmItNDM5NS05OWQ1LTVmOGEzNzg4MWQzNA==",
            com.azure.core.util.Context.NONE);
    }
}
