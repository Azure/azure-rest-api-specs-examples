/** Samples for ServiceTopologies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/servicetopologies_list.json
     */
    /**
     * Sample code: List topologies.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void listTopologies(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager.serviceTopologies().listWithResponse("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
