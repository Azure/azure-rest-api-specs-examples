
/**
 * Samples for NetworkSecurityPerimeterProfiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspProfileDelete.json
     */
    /**
     * Sample code: NspProfilesDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspProfilesDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterProfiles().deleteWithResponse("rg1", "nsp1", "profile1",
            com.azure.core.util.Context.NONE);
    }
}
