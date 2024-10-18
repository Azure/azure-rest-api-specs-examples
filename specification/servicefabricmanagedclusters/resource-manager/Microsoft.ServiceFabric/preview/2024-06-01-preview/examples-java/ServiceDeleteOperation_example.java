
/**
 * Samples for Services Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/
     * examples/ServiceDeleteOperation_example.json
     */
    /**
     * Sample code: Delete a service.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void deleteAService(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.services().delete("resRg", "myCluster", "myApp", "myService", com.azure.core.util.Context.NONE);
    }
}
