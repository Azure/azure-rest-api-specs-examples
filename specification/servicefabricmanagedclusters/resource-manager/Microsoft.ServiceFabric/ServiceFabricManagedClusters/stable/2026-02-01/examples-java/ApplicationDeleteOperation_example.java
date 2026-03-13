
/**
 * Samples for Applications Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ApplicationDeleteOperation_example.json
     */
    /**
     * Sample code: Delete an application.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void deleteAnApplication(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applications().delete("resRg", "myCluster", "myApp", com.azure.core.util.Context.NONE);
    }
}
