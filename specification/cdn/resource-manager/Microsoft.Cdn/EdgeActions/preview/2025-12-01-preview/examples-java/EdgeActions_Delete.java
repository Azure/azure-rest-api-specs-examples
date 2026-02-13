
/**
 * Samples for EdgeActions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/EdgeActions_Delete.json
     */
    /**
     * Sample code: DeleteEdgeAction.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void deleteEdgeAction(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActions().delete("testrg", "edgeAction1", com.azure.core.util.Context.NONE);
    }
}
