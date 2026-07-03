
/**
 * Samples for PublicIpPrefixes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PublicIpPrefixListAll.json
     */
    /**
     * Sample code: List all public IP prefixes.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllPublicIPPrefixes(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPublicIpPrefixes().list(com.azure.core.util.Context.NONE);
    }
}
