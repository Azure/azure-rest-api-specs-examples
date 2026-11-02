
/**
 * Samples for Endpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Endpoints_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Endpoints_Delete_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void endpointsDeleteMaximumSet(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.endpoints().delete("rgprivateTrafficManager", "myProfile", "myEndpoint",
            com.azure.core.util.Context.NONE);
    }
}
