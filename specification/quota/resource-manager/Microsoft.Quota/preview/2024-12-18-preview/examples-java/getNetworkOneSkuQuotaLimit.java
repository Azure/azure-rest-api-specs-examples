
/**
 * Samples for Quota Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2024-12-18-preview/examples/
     * getNetworkOneSkuQuotaLimit.json
     */
    /**
     * Sample code: Quotas_UsagesRequest_ForNetwork.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void quotasUsagesRequestForNetwork(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.quotas().getWithResponse("MinPublicIpInterNetworkPrefixLength",
            "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus",
            com.azure.core.util.Context.NONE);
    }
}
