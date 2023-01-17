/** Samples for ServiceTopologies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/servicetopology_delete.json
     */
    /**
     * Sample code: Delete topology.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void deleteTopology(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .serviceTopologies()
            .deleteByResourceGroupWithResponse("myResourceGroup", "myTopology", com.azure.core.util.Context.NONE);
    }
}
