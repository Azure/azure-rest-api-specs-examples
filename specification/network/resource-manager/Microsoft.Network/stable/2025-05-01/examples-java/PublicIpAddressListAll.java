
/**
 * Samples for PublicIpAddresses List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/PublicIpAddressListAll.json
     */
    /**
     * Sample code: List all public IP addresses.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllPublicIPAddresses(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpAddresses().list(com.azure.core.util.Context.NONE);
    }
}
