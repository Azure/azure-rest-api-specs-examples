
/**
 * Samples for ResourceProvider GetOperationStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/GetOperationStatus.json
     */
    /**
     * Sample code: Get Operation Status.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void getOperationStatus(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.resourceProviders().getOperationStatusWithResponse("HelloWorld", "swaggerExample",
            "YWUzNDFkMzQtZmM5OS00MmUyLWEzNDMtZGJkMDIxZjlmZjgzOzdmYzBiMzhmLTc2NmItNDM5NS05OWQ1LTVmOGEzNzg4MWQzNA==",
            com.azure.core.util.Context.NONE);
    }
}
