
/**
 * Samples for Clusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Clusters_List.json
     */
    /**
     * Sample code: Clusters_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void clustersList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.clusters().list("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
