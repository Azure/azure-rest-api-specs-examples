
/**
 * Samples for VirtualRouters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualRouterDelete.json
     */
    /**
     * Sample code: Delete VirtualRouter.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteVirtualRouter(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualRouters().delete("rg1", "virtualRouter", com.azure.core.util.Context.NONE);
    }
}
