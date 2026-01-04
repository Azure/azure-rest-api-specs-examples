
/**
 * Samples for NetworkSecurityGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkSecurityGroupListAll.
     * json
     */
    /**
     * Sample code: List all network security groups.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllNetworkSecurityGroups(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityGroups().list(com.azure.core.util.Context.NONE);
    }
}
