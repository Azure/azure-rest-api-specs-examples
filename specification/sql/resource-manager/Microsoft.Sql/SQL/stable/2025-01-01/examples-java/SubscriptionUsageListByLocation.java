
/**
 * Samples for SubscriptionUsages ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SubscriptionUsageListByLocation.json
     */
    /**
     * Sample code: List subscription usages in the given location.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listSubscriptionUsagesInTheGivenLocation(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSubscriptionUsages().listByLocation("WestUS", com.azure.core.util.Context.NONE);
    }
}
