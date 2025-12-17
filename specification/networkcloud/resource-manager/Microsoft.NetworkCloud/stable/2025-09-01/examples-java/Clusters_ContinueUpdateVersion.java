
import com.azure.resourcemanager.networkcloud.models.ClusterContinueUpdateVersionMachineGroupTargetingMode;
import com.azure.resourcemanager.networkcloud.models.ClusterContinueUpdateVersionParameters;

/**
 * Samples for Clusters ContinueUpdateVersion.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * Clusters_ContinueUpdateVersion.json
     */
    /**
     * Sample code: Continue update cluster version.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        continueUpdateClusterVersion(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusters().continueUpdateVersion(
            "resourceGroupName", "clusterName", new ClusterContinueUpdateVersionParameters()
                .withMachineGroupTargetingMode(ClusterContinueUpdateVersionMachineGroupTargetingMode.ALPHA_BY_RACK),
            com.azure.core.util.Context.NONE);
    }
}
