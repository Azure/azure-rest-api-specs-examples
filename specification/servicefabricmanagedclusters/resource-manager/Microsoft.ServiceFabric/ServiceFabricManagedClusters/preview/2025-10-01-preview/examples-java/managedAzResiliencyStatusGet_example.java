
/**
 * Samples for ManagedAzResiliencyStatuses Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/managedAzResiliencyStatusGet_example.json
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
