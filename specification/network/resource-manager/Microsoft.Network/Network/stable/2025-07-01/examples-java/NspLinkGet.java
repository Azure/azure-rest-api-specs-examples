
/**
 * Samples for NetworkSecurityPerimeterLinks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspLinkGet.json
     */
    /**
     * Sample code: NspLinksGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspLinksGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterLinks().getWithResponse("rg1", "nsp1", "link1",
            com.azure.core.util.Context.NONE);
    }
}
