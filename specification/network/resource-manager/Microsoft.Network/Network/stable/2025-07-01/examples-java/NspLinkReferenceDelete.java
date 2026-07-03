
/**
 * Samples for NetworkSecurityPerimeterLinkReferences Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspLinkReferenceDelete.json
     */
    /**
     * Sample code: NspLinkReferenceDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspLinkReferenceDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterLinkReferences().delete("rg1", "nsp2", "link1-guid",
            com.azure.core.util.Context.NONE);
    }
}
