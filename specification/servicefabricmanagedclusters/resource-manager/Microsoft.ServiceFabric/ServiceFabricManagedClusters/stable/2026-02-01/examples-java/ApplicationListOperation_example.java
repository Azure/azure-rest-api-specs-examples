
/**
 * Samples for Applications List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ApplicationListOperation_example.json
     */
    /**
     * Sample code: Get a list of application resources.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getAListOfApplicationResources(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applications().list("resRg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
