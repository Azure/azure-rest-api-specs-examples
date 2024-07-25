
/**
 * Samples for ConnectedEnvironments List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * ConnectedEnvironments_ListBySubscription.json
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
