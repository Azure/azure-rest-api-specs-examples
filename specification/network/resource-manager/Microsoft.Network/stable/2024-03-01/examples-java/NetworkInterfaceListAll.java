
/**
 * Samples for NetworkInterfaces List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/NetworkInterfaceListAll.json
     */
    /**
     * Sample code: List all network interfaces.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllNetworkInterfaces(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkInterfaces().list(com.azure.core.util.Context.NONE);
    }
}
