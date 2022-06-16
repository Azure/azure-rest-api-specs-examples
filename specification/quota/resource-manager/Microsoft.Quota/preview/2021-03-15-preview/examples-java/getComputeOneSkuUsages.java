import com.azure.core.util.Context;

/** Samples for Usages Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2021-03-15-preview/examples/getComputeOneSkuUsages.json
     */
    /**
     * Sample code: Quotas_UsagesRequest_ForCompute.
     *
     * @param manager Entry point to QuotaManager.
     */
    public static void quotasUsagesRequestForCompute(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager
            .usages()
            .getWithResponse(
                "standardNDSFamily",
                "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus",
                Context.NONE);
    }
}
