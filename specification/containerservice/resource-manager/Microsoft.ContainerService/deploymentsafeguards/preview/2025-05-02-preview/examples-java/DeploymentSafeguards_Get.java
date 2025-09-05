
/**
 * Samples for DeploymentSafeguards Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-02-preview/DeploymentSafeguards_Get.json
     */
    /**
     * Sample code: Gets a DeploymentSafeguard resource.
     * 
     * @param manager Entry point to ContainerServiceSafeguardsManager.
     */
    public static void getsADeploymentSafeguardResource(
        com.azure.resourcemanager.containerservicesafeguards.ContainerServiceSafeguardsManager manager) {
        manager.deploymentSafeguards().getWithResponse(
            "subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster1",
            com.azure.core.util.Context.NONE);
    }
}
