
/**
 * Samples for PublicIpAddresses Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PublicIpAddressDelete.json
     */
    /**
     * Sample code: Delete public IP address.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deletePublicIPAddress(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpAddresses().delete("rg1", "test-ip", com.azure.core.util.Context.NONE);
    }
}
