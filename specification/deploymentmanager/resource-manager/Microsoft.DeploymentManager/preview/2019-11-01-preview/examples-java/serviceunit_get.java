/** Samples for ServiceUnits Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_get.json
     */
    /**
     * Sample code: Get service unit.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void getServiceUnit(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .serviceUnits()
            .getWithResponse(
                "myResourceGroup", "myTopology", "myService", "myServiceUnit", com.azure.core.util.Context.NONE);
    }
}
