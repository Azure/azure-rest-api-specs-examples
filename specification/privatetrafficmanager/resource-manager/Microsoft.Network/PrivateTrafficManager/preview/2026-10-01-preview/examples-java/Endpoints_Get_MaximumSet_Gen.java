
/**
 * Samples for Endpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Endpoints_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Endpoints_Get_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void
        endpointsGetMaximumSet(com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.endpoints().getWithResponse("rgprivateTrafficManager", "myProfile", "myEndpoint",
            com.azure.core.util.Context.NONE);
    }
}
