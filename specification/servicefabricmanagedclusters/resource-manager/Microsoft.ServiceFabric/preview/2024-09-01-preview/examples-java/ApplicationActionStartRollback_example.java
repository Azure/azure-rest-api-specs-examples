
/**
 * Samples for Applications StartRollback.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/
     * examples/ApplicationActionStartRollback_example.json
     */
    /**
     * Sample code: Start an application upgrade rollback.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void startAnApplicationUpgradeRollback(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applications().startRollback("resRg", "myCluster", "myApp", com.azure.core.util.Context.NONE);
    }
}
