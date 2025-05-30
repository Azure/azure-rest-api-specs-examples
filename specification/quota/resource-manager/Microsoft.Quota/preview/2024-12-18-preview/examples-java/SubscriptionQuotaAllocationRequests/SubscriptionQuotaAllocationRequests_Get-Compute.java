
/**
 * Samples for GroupQuotaSubscriptionAllocationRequest Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2024-12-18-preview/examples/
     * SubscriptionQuotaAllocationRequests/SubscriptionQuotaAllocationRequests_Get-Compute.json
     */
    /**
     * Sample code: SubscriptionQuotaAllocationRequests_Get_Request_ForCompute.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void
        subscriptionQuotaAllocationRequestsGetRequestForCompute(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaSubscriptionAllocationRequests().getWithResponse("E7EC67B3-7657-4966-BFFC-41EFD36BAA09",
            "groupquota1", "Microsoft.Compute", "AE000000-0000-0000-0000-00000000000A",
            com.azure.core.util.Context.NONE);
    }
}
