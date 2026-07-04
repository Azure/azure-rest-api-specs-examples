
/**
 * Samples for CustomIpPrefixes ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/CustomIpPrefixList.json
     */
    /**
     * Sample code: List resource group Custom IP prefixes.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listResourceGroupCustomIPPrefixes(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getCustomIpPrefixes().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
