/** Samples for Steps Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/step_delete.json
     */
    /**
     * Sample code: Delete deployment step.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void deleteDeploymentStep(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .steps()
            .deleteByResourceGroupWithResponse("myResourceGroup", "deploymentStep1", com.azure.core.util.Context.NONE);
    }
}
