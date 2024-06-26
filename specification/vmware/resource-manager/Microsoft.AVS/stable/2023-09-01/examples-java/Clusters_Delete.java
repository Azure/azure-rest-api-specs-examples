
/**
 * Samples for Clusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/Clusters_Delete.json
     */
    /**
     * Sample code: Clusters_Delete.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void clustersDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.clusters().delete("group1", "cloud1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
