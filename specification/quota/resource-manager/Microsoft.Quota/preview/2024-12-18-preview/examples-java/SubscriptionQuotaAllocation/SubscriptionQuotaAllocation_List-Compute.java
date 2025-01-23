
/**
 * Samples for GroupQuotaSubscriptionAllocation List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2024-12-18-preview/examples/
     * SubscriptionQuotaAllocation/SubscriptionQuotaAllocation_List-Compute.json
     */
    /**
     * Sample code: SubscriptionQuotaAllocation_List_ForCompute.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void subscriptionQuotaAllocationListForCompute(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaSubscriptionAllocations().listWithResponse("E7EC67B3-7657-4966-BFFC-41EFD36BAA09",
            "groupquota1", "Microsoft.Compute", "westus", com.azure.core.util.Context.NONE);
    }
}
