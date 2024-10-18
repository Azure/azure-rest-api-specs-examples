
/**
 * Samples for OperationStatus Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/
     * examples/Long_running_operation_status_failed.json
     */
    /**
     * Sample code: Get failed operation status.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getFailedOperationStatus(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.operationStatus().getWithResponse("eastus", "00000000-0000-0000-0000-000000001234",
            com.azure.core.util.Context.NONE);
    }
}
