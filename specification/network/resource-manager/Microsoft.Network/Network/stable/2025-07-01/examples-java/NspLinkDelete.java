
/**
 * Samples for NetworkSecurityPerimeterLinks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspLinkDelete.json
     */
    /**
     * Sample code: NspLinkDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspLinkDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterLinks().delete("rg1", "nsp1", "link1",
            com.azure.core.util.Context.NONE);
    }
}
