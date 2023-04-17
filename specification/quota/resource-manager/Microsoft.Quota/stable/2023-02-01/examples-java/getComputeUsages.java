/** Samples for Usages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/getComputeUsages.json
     */
    /**
     * Sample code: Quotas_listUsagesForCompute.
     *
     * @param manager Entry point to QuotaManager.
     */
    public static void quotasListUsagesForCompute(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager
            .usages()
            .list(
                "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus",
                com.azure.core.util.Context.NONE);
    }
}
