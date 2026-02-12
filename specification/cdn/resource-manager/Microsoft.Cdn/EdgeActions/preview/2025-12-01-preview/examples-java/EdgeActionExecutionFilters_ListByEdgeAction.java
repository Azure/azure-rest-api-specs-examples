
/**
 * Samples for EdgeActionExecutionFilters ListByEdgeAction.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/EdgeActionExecutionFilters_ListByEdgeAction.json
     */
    /**
     * Sample code: ListEdgeActionsExecutionFiltersByEdgeAction.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void
        listEdgeActionsExecutionFiltersByEdgeAction(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActionExecutionFilters().listByEdgeAction("testrg", "edgeAction1",
            com.azure.core.util.Context.NONE);
    }
}
