
/**
 * Samples for VirtualWans ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualWANListByResourceGroup
     * .json
     */
    /**
     * Sample code: VirtualWANListByResourceGroup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualWANListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualWans().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
