
/**
 * Samples for ManagedAzResiliencyStatuses Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/
     * managedAzResiliencyStatusGet_example.json
     */
    /**
     * Sample code: Az Resiliency status of Base Resources.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void azResiliencyStatusOfBaseResources(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedAzResiliencyStatuses().getWithResponse("resourceGroup1", "mycluster1",
            com.azure.core.util.Context.NONE);
    }
}
