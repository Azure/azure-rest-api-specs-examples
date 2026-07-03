
/**
 * Samples for PublicIpAddresses GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PublicIpAddressGet.json
     */
    /**
     * Sample code: Get public IP address.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getPublicIPAddress(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpAddresses().getByResourceGroupWithResponse("rg1", "testDNS-ip", null,
            com.azure.core.util.Context.NONE);
    }
}
