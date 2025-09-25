
/**
 * Samples for GroupQuotaSubscriptionRequests List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/SubscriptionRequests/SubscriptionRequests_List.json
     */
    /**
     * Sample code: GroupQuotaSubscriptionRequests_List.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotaSubscriptionRequestsList(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaSubscriptionRequests().list("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            com.azure.core.util.Context.NONE);
    }
}
