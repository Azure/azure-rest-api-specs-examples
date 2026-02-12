
/**
 * Samples for EdgeActionVersions GetVersionCode.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/EdgeActionVersions_GetVersionCode.json
     */
    /**
     * Sample code: GetVersionCode.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void getVersionCode(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActionVersions().getVersionCode("testrg", "edgeAction1", "version1",
            com.azure.core.util.Context.NONE);
    }
}
