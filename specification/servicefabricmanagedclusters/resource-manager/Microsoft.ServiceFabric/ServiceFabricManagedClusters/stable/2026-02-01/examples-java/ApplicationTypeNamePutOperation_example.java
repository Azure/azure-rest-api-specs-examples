
/**
 * Samples for ApplicationTypes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ApplicationTypeNamePutOperation_example.json
     */
    /**
     * Sample code: Put an application type.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void putAnApplicationType(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applicationTypes().define("myAppType").withExistingManagedCluster("resRg", "myCluster")
            .withRegion("eastus").create();
    }
}
