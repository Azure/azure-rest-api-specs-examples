
/**
 * Samples for EdgeActionExecutionFilters ListByEdgeAction.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeActionExecutionFilters_ListByEdgeAction.json
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
