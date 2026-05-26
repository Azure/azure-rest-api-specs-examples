
/**
 * Samples for OperationStatusResult Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/OperationStatusResultGet.json
     */
    /**
     * Sample code: Get OperationStatusResult.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getOperationStatusResult(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getOperationStatusResults().getWithResponse("rg1", "clustername1",
            "00000000-0000-0000-0000-000000000001", com.azure.core.util.Context.NONE);
    }
}
