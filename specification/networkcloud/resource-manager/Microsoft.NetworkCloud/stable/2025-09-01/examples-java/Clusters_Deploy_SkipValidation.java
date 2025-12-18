
import com.azure.resourcemanager.networkcloud.models.ClusterDeployParameters;
import java.util.Arrays;

/**
 * Samples for Clusters Deploy.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * Clusters_Deploy_SkipValidation.json
     */
    /**
     * Sample code: Deploy cluster skipping validation.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        deployClusterSkippingValidation(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusters().deploy("resourceGroupName", "clusterName",
            new ClusterDeployParameters().withSkipValidationsForMachines(Arrays.asList("bmmName1")),
            com.azure.core.util.Context.NONE);
    }
}
