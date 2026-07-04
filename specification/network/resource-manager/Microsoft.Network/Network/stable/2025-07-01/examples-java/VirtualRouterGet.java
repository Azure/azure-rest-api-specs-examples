
/**
 * Samples for VirtualRouters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualRouterGet.json
     */
    /**
     * Sample code: Get VirtualRouter.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getVirtualRouter(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualRouters().getByResourceGroupWithResponse("rg1", "virtualRouter", null,
            com.azure.core.util.Context.NONE);
    }
}
