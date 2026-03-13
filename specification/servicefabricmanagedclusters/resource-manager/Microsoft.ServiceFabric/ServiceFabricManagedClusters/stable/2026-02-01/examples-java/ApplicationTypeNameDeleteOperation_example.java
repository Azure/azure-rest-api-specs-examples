
/**
 * Samples for ApplicationTypes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ApplicationTypeNameDeleteOperation_example.json
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
