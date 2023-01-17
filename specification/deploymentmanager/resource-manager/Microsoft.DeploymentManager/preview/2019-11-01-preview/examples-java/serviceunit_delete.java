/** Samples for ServiceUnits Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_delete.json
     */
    /**
     * Sample code: Delete service unit.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void deleteServiceUnit(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .serviceUnits()
            .deleteWithResponse(
                "myResourceGroup", "myTopology", "myService", "myServiceUnit", com.azure.core.util.Context.NONE);
    }
}
