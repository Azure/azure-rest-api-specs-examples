
/**
 * Samples for ConnectedEnvironments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/ConnectedEnvironments_Delete
     * .json
     */
    /**
     * Sample code: Delete connected environment by connectedEnvironmentName.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteConnectedEnvironmentByConnectedEnvironmentName(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironments().delete("examplerg", "examplekenv", com.azure.core.util.Context.NONE);
    }
}
