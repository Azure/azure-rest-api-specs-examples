import com.azure.core.util.Context;

/** Samples for Quota List. */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2021-03-15-preview/examples/getNetworkQuotaLimits.json
     */
    /**
     * Sample code: Quotas_listQuotaLimitsForNetwork.
     *
     * @param manager Entry point to QuotaManager.
     */
    public static void quotasListQuotaLimitsForNetwork(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager
            .quotas()
            .list(
                "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus",
                Context.NONE);
    }
}
