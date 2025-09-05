
/**
 * Samples for DeploymentSafeguards List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-02-preview/DeploymentSafeguards_List.json
     */
    /**
     * Sample code: Lists DeploymentSafeguards by parent resource.
     * 
     * @param manager Entry point to ContainerServiceSafeguardsManager.
     */
    public static void listsDeploymentSafeguardsByParentResource(
        com.azure.resourcemanager.containerservicesafeguards.ContainerServiceSafeguardsManager manager) {
        manager.deploymentSafeguards().list(
            "subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster1",
            com.azure.core.util.Context.NONE);
    }
}
