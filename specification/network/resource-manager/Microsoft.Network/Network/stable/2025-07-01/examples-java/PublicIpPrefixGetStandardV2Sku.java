
/**
 * Samples for PublicIpPrefixes GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PublicIpPrefixGetStandardV2Sku.json
     */
    /**
     * Sample code: Get public IP prefix with StandardV2 sku.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getPublicIPPrefixWithStandardV2Sku(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpPrefixes().getByResourceGroupWithResponse("rg1", "test-ipprefix", null,
            com.azure.core.util.Context.NONE);
    }
}
