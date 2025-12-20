
/**
 * Samples for Addons Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Addons_Get_HCX.json
     */
    /**
     * Sample code: Addons_Get_HCX.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void addonsGetHCX(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.addons().getWithResponse("group1", "cloud1", "hcx", com.azure.core.util.Context.NONE);
    }
}
