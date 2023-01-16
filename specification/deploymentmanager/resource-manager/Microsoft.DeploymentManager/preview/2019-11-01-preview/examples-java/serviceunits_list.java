/** Samples for ServiceUnits List. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunits_list.json
     */
    /**
     * Sample code: List service units.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void listServiceUnits(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .serviceUnits()
            .listWithResponse("myResourceGroup", "myTopology", "myService", com.azure.core.util.Context.NONE);
    }
}
