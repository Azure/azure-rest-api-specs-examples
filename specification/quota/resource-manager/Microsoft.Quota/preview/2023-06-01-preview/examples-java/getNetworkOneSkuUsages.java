
/**
 * Samples for Usages Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/getNetworkOneSkuUsages.
     * json
     */
    /**
     * Sample code: Quotas_UsagesRequest_ForNetwork.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void quotasUsagesRequestForNetwork(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.usages().getWithResponse("MinPublicIpInterNetworkPrefixLength",
            "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus",
            com.azure.core.util.Context.NONE);
    }
}
