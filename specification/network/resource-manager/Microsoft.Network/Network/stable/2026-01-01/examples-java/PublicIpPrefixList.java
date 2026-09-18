
/**
 * Samples for PublicIpPrefixes ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/PublicIpPrefixList.json
     */
    /**
     * Sample code: List resource group public IP prefixes.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listResourceGroupPublicIPPrefixes(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpPrefixes().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
