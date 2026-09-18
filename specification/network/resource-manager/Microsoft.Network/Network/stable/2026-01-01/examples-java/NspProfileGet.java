
/**
 * Samples for NetworkSecurityPerimeterProfiles Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NspProfileGet.json
     */
    /**
     * Sample code: NspProfilesGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspProfilesGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterProfiles().getWithResponse("rg1", "nsp1", "profile1",
            com.azure.core.util.Context.NONE);
    }
}
