
/**
 * Samples for CustomIpPrefixes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/CustomIpPrefixListAll.json
     */
    /**
     * Sample code: List all custom IP prefixes.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllCustomIPPrefixes(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getCustomIpPrefixes().list(com.azure.core.util.Context.NONE);
    }
}
