
/**
 * Samples for EdgeActionVersions ListByEdgeAction.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeActionVersions_ListByEdgeAction.json
     */
    /**
     * Sample code: ListEdgeActionVersionsByEdgeAction.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void
        listEdgeActionVersionsByEdgeAction(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActionVersions().listByEdgeAction("testrg", "edgeAction1", com.azure.core.util.Context.NONE);
    }
}
