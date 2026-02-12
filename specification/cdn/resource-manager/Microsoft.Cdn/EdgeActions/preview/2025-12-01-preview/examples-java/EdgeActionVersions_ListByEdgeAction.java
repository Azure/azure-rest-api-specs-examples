
/**
 * Samples for EdgeActionVersions ListByEdgeAction.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/EdgeActionVersions_ListByEdgeAction.json
     */
    /**
     * Sample code: GetEdgeActionVersionsByEdgeAction.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void
        getEdgeActionVersionsByEdgeAction(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActionVersions().listByEdgeAction("testrg", "edgeAction1", com.azure.core.util.Context.NONE);
    }
}
