import com.azure.core.util.Context;

/** Samples for ConnectedEnvironments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ConnectedEnvironments_Delete.json
     */
    /**
     * Sample code: Delete connected environment by connectedEnvironmentName.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteConnectedEnvironmentByConnectedEnvironmentName(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironments().delete("examplerg", "examplekenv", Context.NONE);
    }
}
