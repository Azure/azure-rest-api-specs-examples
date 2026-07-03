
/**
 * Samples for NetworkSecurityPerimeterLinkReferences Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspLinkReferenceGet.json
     */
    /**
     * Sample code: NspLinkReferencesGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspLinkReferencesGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterLinkReferences().getWithResponse("rg1", "nsp2", "link1-guid",
            com.azure.core.util.Context.NONE);
    }
}
