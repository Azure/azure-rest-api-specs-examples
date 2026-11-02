
/**
 * Samples for Sites Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Sites_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Sites_Get_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void
        sitesGetMaximumSet(com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.sites().getWithResponse("rgprivateTrafficManager", "myTopologyMap", "mySite",
            com.azure.core.util.Context.NONE);
    }
}
