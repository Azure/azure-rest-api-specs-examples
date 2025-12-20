
/**
 * Samples for Addons Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Addons_Get_VR.json
     */
    /**
     * Sample code: Addons_Get_VR.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void addonsGetVR(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.addons().getWithResponse("group1", "cloud1", "vr", com.azure.core.util.Context.NONE);
    }
}
