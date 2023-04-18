/** Samples for Quota Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/getComputeOneSkuQuotaLimit.json
     */
    /**
     * Sample code: Quotas_Get_Request_ForCompute.
     *
     * @param manager Entry point to QuotaManager.
     */
    public static void quotasGetRequestForCompute(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager
            .quotas()
            .getWithResponse(
                "standardNDSFamily",
                "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus",
                com.azure.core.util.Context.NONE);
    }
}
