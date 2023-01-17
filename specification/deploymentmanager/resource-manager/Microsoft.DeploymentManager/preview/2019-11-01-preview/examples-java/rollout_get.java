/** Samples for Rollouts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollout_get.json
     */
    /**
     * Sample code: Get rollout.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void getRollout(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .rollouts()
            .getByResourceGroupWithResponse("myResourceGroup", "myRollout", null, com.azure.core.util.Context.NONE);
    }
}
