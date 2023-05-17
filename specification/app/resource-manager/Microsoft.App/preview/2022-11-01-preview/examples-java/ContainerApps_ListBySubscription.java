/** Samples for ContainerApps List. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/ContainerApps_ListBySubscription.json
     */
    /**
     * Sample code: List Container Apps by subscription.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listContainerAppsBySubscription(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().list(com.azure.core.util.Context.NONE);
    }
}
