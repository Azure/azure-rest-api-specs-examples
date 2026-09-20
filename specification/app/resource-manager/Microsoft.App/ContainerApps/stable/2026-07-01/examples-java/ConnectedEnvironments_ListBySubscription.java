
/**
 * Samples for ConnectedEnvironments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ConnectedEnvironments_ListBySubscription.json
     */
    /**
     * Sample code: List connected environments by subscription.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listConnectedEnvironmentsBySubscription(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironments().list(com.azure.core.util.Context.NONE);
    }
}
