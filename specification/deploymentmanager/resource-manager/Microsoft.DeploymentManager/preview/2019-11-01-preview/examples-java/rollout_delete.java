/** Samples for Rollouts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollout_delete.json
     */
    /**
     * Sample code: Delete rollout.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void deleteRollout(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .rollouts()
            .deleteByResourceGroupWithResponse("myResourceGroup", "myRollout", com.azure.core.util.Context.NONE);
    }
}
