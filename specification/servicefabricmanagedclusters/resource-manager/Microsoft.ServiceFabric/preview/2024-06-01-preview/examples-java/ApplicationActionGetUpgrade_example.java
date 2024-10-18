
/**
 * Samples for Applications ReadUpgrade.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/
     * examples/ApplicationActionGetUpgrade_example.json
     */
    /**
     * Sample code: Get an application upgrade.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getAnApplicationUpgrade(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applications().readUpgrade("resRg", "myCluster", "myApp", com.azure.core.util.Context.NONE);
    }
}
