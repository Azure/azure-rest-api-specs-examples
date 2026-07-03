
/**
 * Samples for PublicIpAddresses ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PublicIpAddressList.json
     */
    /**
     * Sample code: List resource group public IP addresses.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listResourceGroupPublicIPAddresses(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpAddresses().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
