
/**
 * Samples for EdgeActionVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeActionVersions_Get.json
     */
    /**
     * Sample code: GetEdgeActionVersion.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void getEdgeActionVersion(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActionVersions().getWithResponse("testrg", "edgeAction1", "version1",
            com.azure.core.util.Context.NONE);
    }
}
