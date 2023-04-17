/** Samples for Usages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/getNetworkUsages.json
     */
    /**
     * Sample code: Quotas_listUsagesForNetwork.
     *
     * @param manager Entry point to QuotaManager.
     */
    public static void quotasListUsagesForNetwork(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager
            .usages()
            .list(
                "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus",
                com.azure.core.util.Context.NONE);
    }
}
