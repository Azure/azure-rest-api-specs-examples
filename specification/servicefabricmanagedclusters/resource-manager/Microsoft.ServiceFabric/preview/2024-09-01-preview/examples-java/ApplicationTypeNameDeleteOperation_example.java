
/**
 * Samples for ApplicationTypes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/
     * examples/ApplicationTypeNameDeleteOperation_example.json
     */
    /**
     * Sample code: Delete an application type.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void deleteAnApplicationType(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applicationTypes().delete("resRg", "myCluster", "myAppType", com.azure.core.util.Context.NONE);
    }
}
