
/**
 * Samples for Jobs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/Jobs_ListBySubscription.json
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
