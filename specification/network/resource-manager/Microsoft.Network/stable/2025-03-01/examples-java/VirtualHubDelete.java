
/**
 * Samples for VirtualHubs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualHubDelete.json
     */
    /**
     * Sample code: VirtualHubDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualHubDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualHubs().delete("rg1", "virtualHub1",
            com.azure.core.util.Context.NONE);
    }
}
