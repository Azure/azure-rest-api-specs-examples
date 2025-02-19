
/**
 * Samples for ApplicationTypes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/
     * examples/ApplicationTypeNamePutOperation_example.json
     */
    /**
     * Sample code: Put an application type.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void putAnApplicationType(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applicationTypes().define("myAppType").withExistingManagedcluster("resRg", "myCluster")
            .withRegion("eastus").create();
    }
}
