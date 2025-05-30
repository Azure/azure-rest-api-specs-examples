
import com.azure.resourcemanager.quota.fluent.models.SubscriptionQuotaAllocationsListInner;
import com.azure.resourcemanager.quota.models.SubscriptionQuotaAllocations;
import com.azure.resourcemanager.quota.models.SubscriptionQuotaAllocationsListProperties;
import com.azure.resourcemanager.quota.models.SubscriptionQuotaAllocationsProperties;
import java.util.Arrays;

/**
 * Samples for GroupQuotaSubscriptionAllocationRequest Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2025-03-01/examples/
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
            "Microsoft.Compute", "westus",
            new SubscriptionQuotaAllocationsListInner()
                .withProperties(new SubscriptionQuotaAllocationsListProperties().withValue(Arrays.asList(
                    new SubscriptionQuotaAllocations().withProperties(new SubscriptionQuotaAllocationsProperties()
                        .withResourceName("standardddv4family").withLimit(110L)),
                    new SubscriptionQuotaAllocations().withProperties(new SubscriptionQuotaAllocationsProperties()
                        .withResourceName("standardav2family").withLimit(110L))))),
            com.azure.core.util.Context.NONE);
    }
}
