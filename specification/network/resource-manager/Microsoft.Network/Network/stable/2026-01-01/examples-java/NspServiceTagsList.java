
/**
 * Samples for NetworkSecurityPerimeterServiceTags List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NspServiceTagsList.json
     */
    /**
     * Sample code: NSPServiceTagsList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nSPServiceTagsList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterServiceTags().list("westus",
            com.azure.core.util.Context.NONE);
    }
}
