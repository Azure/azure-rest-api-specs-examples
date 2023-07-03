import com.azure.resourcemanager.networkcloud.models.ClusterDeployParameters;

/** Samples for Clusters Deploy. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/Clusters_Deploy.json
     */
    /**
     * Sample code: Deploy cluster.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deployCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .clusters()
            .deploy(
                "resourceGroupName", "clusterName", new ClusterDeployParameters(), com.azure.core.util.Context.NONE);
    }
}
