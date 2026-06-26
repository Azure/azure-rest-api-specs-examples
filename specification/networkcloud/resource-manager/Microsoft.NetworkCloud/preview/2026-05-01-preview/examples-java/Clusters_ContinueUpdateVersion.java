
import com.azure.resourcemanager.networkcloud.models.ClusterContinueUpdateVersionMachineGroupTargetingMode;
import com.azure.resourcemanager.networkcloud.models.ClusterContinueUpdateVersionParameters;
import com.azure.resourcemanager.networkcloud.models.ClusterContinueUpdateVersionSafeguardMode;

/**
 * Samples for Clusters ContinueUpdateVersion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Clusters_ContinueUpdateVersion.json
     */
    /**
     * Sample code: Continue update cluster version.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        continueUpdateClusterVersion(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusters().continueUpdateVersion("resourceGroupName", "clusterName",
            new ClusterContinueUpdateVersionParameters()
                .withMachineGroupTargetingMode(ClusterContinueUpdateVersionMachineGroupTargetingMode.ALPHA_BY_RACK)
                .withSafeguardMode(ClusterContinueUpdateVersionSafeguardMode.ALL),
            com.azure.core.util.Context.NONE);
    }
}
