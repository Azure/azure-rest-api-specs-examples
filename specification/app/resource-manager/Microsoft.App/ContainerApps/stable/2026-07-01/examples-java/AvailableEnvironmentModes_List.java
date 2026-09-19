
/**
 * Samples for AvailableEnvironmentModes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/AvailableEnvironmentModes_List.json
     */
    /**
     * Sample code: AvailableEnvironmentModes_List.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        availableEnvironmentModesList(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.availableEnvironmentModes().list("East US", com.azure.core.util.Context.NONE);
    }
}
