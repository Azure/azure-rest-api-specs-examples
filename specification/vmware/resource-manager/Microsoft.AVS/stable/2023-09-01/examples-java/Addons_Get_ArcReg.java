
/**
 * Samples for Addons Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/Addons_Get_ArcReg.json
     */
    /**
     * Sample code: Addons_Get_ArcReg.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void addonsGetArcReg(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.addons().getWithResponse("group1", "cloud1", "arc", com.azure.core.util.Context.NONE);
    }
}
