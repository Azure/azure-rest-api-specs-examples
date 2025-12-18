
import com.azure.resourcemanager.networkcloud.models.ClusterScanRuntimeParameters;
import com.azure.resourcemanager.networkcloud.models.ClusterScanRuntimeParametersScanActivity;

/**
 * Samples for Clusters ScanRuntime.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * Clusters_ScanRuntime.json
     */
    /**
     * Sample code: Execute a runtime protection scan on the cluster.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        executeARuntimeProtectionScanOnTheCluster(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusters().scanRuntime("resourceGroupName", "clusterName",
            new ClusterScanRuntimeParameters().withScanActivity(ClusterScanRuntimeParametersScanActivity.SCAN),
            com.azure.core.util.Context.NONE);
    }
}
