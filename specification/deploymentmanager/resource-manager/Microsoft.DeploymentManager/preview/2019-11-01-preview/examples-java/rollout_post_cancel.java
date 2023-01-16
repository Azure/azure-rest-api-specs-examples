/** Samples for Rollouts Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollout_post_cancel.json
     */
    /**
     * Sample code: Cancel rollout.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void cancelRollout(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager.rollouts().cancelWithResponse("myResourceGroup", "myRollout", com.azure.core.util.Context.NONE);
    }
}
