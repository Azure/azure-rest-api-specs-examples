
/**
 * Samples for OperationStatusResult List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/OperationStatusResultList.json
     */
    /**
     * Sample code: List of OperationStatusResult.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        listOfOperationStatusResult(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getOperationStatusResults().list("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
