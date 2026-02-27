
/**
 * Samples for VirtualRouters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualRouterListByResourceGroup.json
     */
    /**
     * Sample code: List all Virtual Router for a given resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAllVirtualRouterForAGivenResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualRouters().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
