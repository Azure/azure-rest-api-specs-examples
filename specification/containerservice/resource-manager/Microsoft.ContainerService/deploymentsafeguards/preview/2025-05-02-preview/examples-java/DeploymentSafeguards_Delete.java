
/**
 * Samples for DeploymentSafeguards Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-02-preview/DeploymentSafeguards_Delete.json
     */
    /**
     * Sample code: Deletes a DeploymentSafeguard resource asynchronously with a long running operation.
     * 
     * @param manager Entry point to ContainerServiceSafeguardsManager.
     */
    public static void deletesADeploymentSafeguardResourceAsynchronouslyWithALongRunningOperation(
        com.azure.resourcemanager.containerservicesafeguards.ContainerServiceSafeguardsManager manager) {
        manager.deploymentSafeguards().delete(
            "subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster1",
            com.azure.core.util.Context.NONE);
    }
}
