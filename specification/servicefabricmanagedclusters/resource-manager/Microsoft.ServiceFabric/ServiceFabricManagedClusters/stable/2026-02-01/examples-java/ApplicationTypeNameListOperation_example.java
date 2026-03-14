
/**
 * Samples for ApplicationTypes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ApplicationTypeNameListOperation_example.json
     */
    /**
     * Sample code: Get a list of application type name resources.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getAListOfApplicationTypeNameResources(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applicationTypes().list("resRg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
