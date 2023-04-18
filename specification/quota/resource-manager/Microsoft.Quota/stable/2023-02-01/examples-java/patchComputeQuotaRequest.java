import com.azure.resourcemanager.quota.models.CurrentQuotaLimitBase;
import com.azure.resourcemanager.quota.models.LimitObject;
import com.azure.resourcemanager.quota.models.QuotaProperties;
import com.azure.resourcemanager.quota.models.ResourceName;

/** Samples for Quota Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/patchComputeQuotaRequest.json
     */
    /**
     * Sample code: Quotas_Request_PatchForCompute.
     *
     * @param manager Entry point to QuotaManager.
     */
    public static void quotasRequestPatchForCompute(com.azure.resourcemanager.quota.QuotaManager manager) {
        CurrentQuotaLimitBase resource =
            manager
                .quotas()
                .getWithResponse(
                    "standardFSv2Family",
                    "subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Compute/locations/eastus",
                    com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withProperties(
                new QuotaProperties()
                    .withLimit(new LimitObject().withValue(10))
                    .withName(new ResourceName().withValue("standardFSv2Family")))
            .apply();
    }
}
