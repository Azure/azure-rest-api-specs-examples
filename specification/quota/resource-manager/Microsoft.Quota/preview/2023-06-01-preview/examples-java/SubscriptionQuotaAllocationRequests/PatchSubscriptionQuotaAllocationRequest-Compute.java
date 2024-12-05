
import com.azure.resourcemanager.quota.fluent.models.QuotaAllocationRequestStatusInner;
import com.azure.resourcemanager.quota.models.QuotaAllocationRequestBase;

/**
 * Samples for GroupQuotaSubscriptionAllocationRequest Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/
     * SubscriptionQuotaAllocationRequests/PatchSubscriptionQuotaAllocationRequest-Compute.json
     */
    /**
     * Sample code: SubscriptionQuotaAllocation_Patch_Request_ForCompute.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void
        subscriptionQuotaAllocationPatchRequestForCompute(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.groupQuotaSubscriptionAllocationRequests().update("E7EC67B3-7657-4966-BFFC-41EFD36BAA09", "groupquota1",
            "Microsoft.Compute", "standardav2family",
            new QuotaAllocationRequestStatusInner()
                .withRequestedResource(new QuotaAllocationRequestBase().withLimit(10L).withRegion("westus")),
            com.azure.core.util.Context.NONE);
    }
}
