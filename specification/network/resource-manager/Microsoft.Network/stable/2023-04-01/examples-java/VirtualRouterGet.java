/** Samples for VirtualRouters GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualRouterGet.json
     */
    /**
     * Sample code: Get VirtualRouter.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualRouter(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualRouters()
            .getByResourceGroupWithResponse("rg1", "virtualRouter", null, com.azure.core.util.Context.NONE);
    }
}
