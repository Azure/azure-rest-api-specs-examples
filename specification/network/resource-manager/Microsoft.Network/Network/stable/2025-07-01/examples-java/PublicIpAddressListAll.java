
/**
 * Samples for PublicIpAddresses List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PublicIpAddressListAll.json
     */
    /**
     * Sample code: List all public IP addresses.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllPublicIPAddresses(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpAddresses().list(com.azure.core.util.Context.NONE);
    }
}
