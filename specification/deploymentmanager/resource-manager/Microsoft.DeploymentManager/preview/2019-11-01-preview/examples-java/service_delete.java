/** Samples for Services Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/service_delete.json
     */
    /**
     * Sample code: Delete service.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void deleteService(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .services()
            .deleteWithResponse("myResourceGroup", "myTopology", "myService", com.azure.core.util.Context.NONE);
    }
}
