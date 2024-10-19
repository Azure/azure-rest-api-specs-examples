
/**
 * Samples for ApplicationTypeVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/
     * examples/ApplicationTypeVersionDeleteOperation_example.json
     */
    /**
     * Sample code: Delete an application type version.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void deleteAnApplicationTypeVersion(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applicationTypeVersions().delete("resRg", "myCluster", "myAppType", "1.0",
            com.azure.core.util.Context.NONE);
    }
}
