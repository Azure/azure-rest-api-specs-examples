/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/operations_list.json
     */
    /**
     * Sample code: Get operations.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void getOperations(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager.operations().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
