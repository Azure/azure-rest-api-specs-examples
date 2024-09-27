
/**
 * Samples for NetworkProfiles List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/NetworkProfileListAll.json
     */
    /**
     * Sample code: List all network profiles.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllNetworkProfiles(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkProfiles().list(com.azure.core.util.Context.NONE);
    }
}
