
/**
 * Samples for EdgeActionVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeActionVersions_Delete.json
     */
    /**
     * Sample code: DeleteEdgeActionVersion.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void deleteEdgeActionVersion(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActionVersions().delete("testrg", "edgeAction1", "version1", com.azure.core.util.Context.NONE);
    }
}
