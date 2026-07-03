
/**
 * Samples for PublicIpAddresses GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PublicIpAddressGetStandardV2Sku.json
     */
    /**
     * Sample code: Get public IP address with StandardV2 sku.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getPublicIPAddressWithStandardV2Sku(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpAddresses().getByResourceGroupWithResponse("rg1", "testDNS-ip", null,
            com.azure.core.util.Context.NONE);
    }
}
