
/**
 * Samples for NetworkSecurityPerimeterLinkReferences List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspLinkReferenceList.json
     */
    /**
     * Sample code: NspLinkReferenceList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspLinkReferenceList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterLinkReferences().list("rg1", "nsp2", null, null,
            com.azure.core.util.Context.NONE);
    }
}
