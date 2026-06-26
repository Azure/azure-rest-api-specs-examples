
import com.azure.resourcemanager.networkcloud.models.ClusterUpdateVersionParameters;
import com.azure.resourcemanager.networkcloud.models.ClusterUpdateVersionSafeguardMode;

/**
 * Samples for Clusters UpdateVersion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Clusters_UpdateVersion.json
     */
    /**
     * Sample code: Update cluster version.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void updateClusterVersion(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusters().updateVersion(
            "resourceGroupName", "clusterName", new ClusterUpdateVersionParameters()
                .withSafeguardMode(ClusterUpdateVersionSafeguardMode.ALL).withTargetClusterVersion("2.0"),
            com.azure.core.util.Context.NONE);
    }
}
