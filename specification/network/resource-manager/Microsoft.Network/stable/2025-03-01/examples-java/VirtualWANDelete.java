
/**
 * Samples for VirtualWans Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualWANDelete.json
     */
    /**
     * Sample code: VirtualWANDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualWANDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualWans().delete("rg1", "virtualWan1",
            com.azure.core.util.Context.NONE);
    }
}
