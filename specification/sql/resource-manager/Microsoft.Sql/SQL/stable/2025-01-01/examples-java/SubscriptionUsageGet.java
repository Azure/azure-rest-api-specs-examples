
/**
 * Samples for SubscriptionUsages Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SubscriptionUsageGet.json
     */
    /**
     * Sample code: Get specific subscription usage in the given location.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getSpecificSubscriptionUsageInTheGivenLocation(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSubscriptionUsages().getWithResponse("WestUS", "ServerQuota",
            com.azure.core.util.Context.NONE);
    }
}
