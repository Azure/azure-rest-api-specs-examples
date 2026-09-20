
/**
 * Samples for ContainerAppsSessionPools List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SessionPools_ListBySubscription.json
     */
    /**
     * Sample code: List Session Pools by subscription.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listSessionPoolsBySubscription(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsSessionPools().list(com.azure.core.util.Context.NONE);
    }
}
