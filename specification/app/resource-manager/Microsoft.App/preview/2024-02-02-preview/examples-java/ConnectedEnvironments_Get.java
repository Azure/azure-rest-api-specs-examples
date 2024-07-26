
/**
 * Samples for ConnectedEnvironments GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/ConnectedEnvironments_Get.
     * json
     */
    /**
     * Sample code: Get connected environment by connectedEnvironmentName.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getConnectedEnvironmentByConnectedEnvironmentName(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironments().getByResourceGroupWithResponse("examplerg", "examplekenv",
            com.azure.core.util.Context.NONE);
    }
}
