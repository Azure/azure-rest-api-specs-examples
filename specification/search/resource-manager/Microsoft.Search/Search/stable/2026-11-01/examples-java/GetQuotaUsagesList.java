
/**
 * Samples for Usages ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/GetQuotaUsagesList.json
     */
    /**
     * Sample code: GetQuotaUsagesList.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void getQuotaUsagesList(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getUsages().listBySubscription("westus", com.azure.core.util.Context.NONE);
    }
}
