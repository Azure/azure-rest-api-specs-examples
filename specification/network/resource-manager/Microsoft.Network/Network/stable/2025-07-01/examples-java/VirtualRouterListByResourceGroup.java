
/**
 * Samples for VirtualRouters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualRouterListByResourceGroup.json
     */
    /**
     * Sample code: List all Virtual Router for a given resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listAllVirtualRouterForAGivenResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualRouters().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
