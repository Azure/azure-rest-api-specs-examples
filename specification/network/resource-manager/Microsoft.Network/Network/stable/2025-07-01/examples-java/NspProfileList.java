
/**
 * Samples for NetworkSecurityPerimeterProfiles List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspProfileList.json
     */
    /**
     * Sample code: NspProfilesList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspProfilesList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterProfiles().list("rg1", "nsp1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
