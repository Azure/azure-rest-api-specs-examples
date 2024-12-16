
/**
 * Samples for ApplicationTypeVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/
     * examples/ApplicationTypeVersionPutOperation_example.json
     */
    /**
     * Sample code: Put an application type version.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void putAnApplicationTypeVersion(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applicationTypeVersions().define("1.0").withExistingApplicationType("resRg", "myCluster", "myAppType")
            .withRegion("eastus").withAppPackageUrl("http://fakelink.test.com/MyAppType").create();
    }
}
