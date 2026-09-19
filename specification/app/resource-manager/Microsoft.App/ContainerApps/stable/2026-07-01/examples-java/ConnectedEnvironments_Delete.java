
/**
 * Samples for ConnectedEnvironments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ConnectedEnvironments_Delete.json
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
