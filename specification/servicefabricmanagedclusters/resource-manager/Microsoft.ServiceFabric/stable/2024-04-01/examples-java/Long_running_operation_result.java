
/**
 * Samples for OperationResults Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/
     * Long_running_operation_result.json
     */
    /**
     * Sample code: Get operation result.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getOperationResult(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.operationResults().getWithResponse("eastus", "00000000-0000-0000-0000-000000001234",
            com.azure.core.util.Context.NONE);
    }
}
