
/**
 * Samples for GroupQuotaSubscriptionAllocationRequest List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/
     * SubscriptionQuotaAllocationRequests/SubscriptionQuotaAllocationRequests_List-Compute.json
     */
    /**
     * Sample code: SubscriptionQuotaAllocation_List_Request_ForCompute.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void
        subscriptionQuotaAllocationListRequestForCompute(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaSubscriptionAllocationRequests().list("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            "Microsoft.Compute", "location eq westus", com.azure.core.util.Context.NONE);
    }
}
