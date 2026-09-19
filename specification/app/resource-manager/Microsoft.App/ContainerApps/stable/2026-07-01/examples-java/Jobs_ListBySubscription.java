
/**
 * Samples for Jobs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Jobs_ListBySubscription.json
     */
    /**
     * Sample code: List Container Apps Jobs by subscription.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listContainerAppsJobsBySubscription(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().list(com.azure.core.util.Context.NONE);
    }
}
