
/**
 * Samples for GroupQuotaSubscriptionRequests Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/SubscriptionRequests/SubscriptionRequests_Get.json
     */
    /**
     * Sample code: GroupQuotaSubscriptionRequests_Get.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void groupQuotaSubscriptionRequestsGet(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaSubscriptionRequests().getWithResponse("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            "00000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
