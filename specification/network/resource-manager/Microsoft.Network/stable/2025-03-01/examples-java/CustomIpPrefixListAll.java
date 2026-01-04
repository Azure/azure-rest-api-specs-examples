
/**
 * Samples for CustomIpPrefixes List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/CustomIpPrefixListAll.json
     */
    /**
     * Sample code: List all custom IP prefixes.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllCustomIPPrefixes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getCustomIpPrefixes().list(com.azure.core.util.Context.NONE);
    }
}
