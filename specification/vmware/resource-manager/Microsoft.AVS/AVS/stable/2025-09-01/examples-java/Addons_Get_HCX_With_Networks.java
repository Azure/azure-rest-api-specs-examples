
/**
 * Samples for Addons Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Addons_Get_HCX_With_Networks.json
     */
    /**
     * Sample code: Addons_Get_HCX_With_Networks.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void addonsGetHCXWithNetworks(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.addons().getWithResponse("group1", "cloud1", "hcx", com.azure.core.util.Context.NONE);
    }
}
