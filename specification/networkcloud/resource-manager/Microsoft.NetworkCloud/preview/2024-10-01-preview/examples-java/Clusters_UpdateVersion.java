
import com.azure.resourcemanager.networkcloud.models.ClusterUpdateVersionParameters;

/**
 * Samples for Clusters UpdateVersion.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-10-01-preview/examples/
     * Clusters_UpdateVersion.json
     */
    /**
     * Sample code: Update cluster version.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void updateClusterVersion(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusters().updateVersion("resourceGroupName", "clusterName",
            new ClusterUpdateVersionParameters().withTargetClusterVersion("2.0"), com.azure.core.util.Context.NONE);
    }
}
