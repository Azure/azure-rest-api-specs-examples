
/**
 * Samples for Sites Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Sites_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Sites_Delete_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void
        sitesDeleteMaximumSet(com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.sites().delete("rgprivateTrafficManager", "myTopologyMap", "mySite", com.azure.core.util.Context.NONE);
    }
}
