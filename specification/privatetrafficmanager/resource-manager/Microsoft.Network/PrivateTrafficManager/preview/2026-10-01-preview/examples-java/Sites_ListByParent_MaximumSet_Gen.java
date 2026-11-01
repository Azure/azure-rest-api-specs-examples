
/**
 * Samples for Sites ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Sites_ListByParent_MaximumSet_Gen.json
     */
    /**
     * Sample code: Sites_ListByParent_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void sitesListByParentMaximumSet(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.sites().listByParent("rgprivateTrafficManager", "myTopologyMap", com.azure.core.util.Context.NONE);
    }
}
