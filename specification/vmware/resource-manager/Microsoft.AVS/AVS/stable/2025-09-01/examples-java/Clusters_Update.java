
import com.azure.resourcemanager.avs.models.Cluster;

/**
 * Samples for Clusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Clusters_Update.json
     */
    /**
     * Sample code: Clusters_Update.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void clustersUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        Cluster resource = manager.clusters()
            .getWithResponse("group1", "cloud1", "cluster1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withClusterSize(4).apply();
    }
}
