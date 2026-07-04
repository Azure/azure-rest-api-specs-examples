
/**
 * Samples for VirtualWans ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualWANListByResourceGroup.json
     */
    /**
     * Sample code: VirtualWANListByResourceGroup.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualWANListByResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualWans().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
