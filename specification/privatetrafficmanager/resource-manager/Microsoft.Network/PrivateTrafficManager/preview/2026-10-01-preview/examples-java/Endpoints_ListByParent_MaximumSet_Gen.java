
/**
 * Samples for Endpoints ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Endpoints_ListByParent_MaximumSet_Gen.json
     */
    /**
     * Sample code: Endpoints_ListByParent_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void endpointsListByParentMaximumSet(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.endpoints().listByParent("rgprivateTrafficManager", "myProfile", com.azure.core.util.Context.NONE);
    }
}
