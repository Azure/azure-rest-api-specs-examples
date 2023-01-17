/** Samples for Services List. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/services_list.json
     */
    /**
     * Sample code: List services.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void listServices(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager.services().listWithResponse("myResourceGroup", "myTopology", com.azure.core.util.Context.NONE);
    }
}
