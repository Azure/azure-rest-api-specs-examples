/** Samples for Services Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/service_get.json
     */
    /**
     * Sample code: Get service.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void getService(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .services()
            .getWithResponse("myResourceGroup", "myTopology", "myService", com.azure.core.util.Context.NONE);
    }
}
