
/**
 * Samples for ContainerApps List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerApps_ListBySubscription.json
     */
    /**
     * Sample code: List Container Apps by subscription.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listContainerAppsBySubscription(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().list(com.azure.core.util.Context.NONE);
    }
}
