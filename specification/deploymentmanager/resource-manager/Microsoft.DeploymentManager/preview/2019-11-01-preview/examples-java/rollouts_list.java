/** Samples for Rollouts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollouts_list.json
     */
    /**
     * Sample code: List rollouts by resource group.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void listRolloutsByResourceGroup(
        com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager.rollouts().listWithResponse("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
