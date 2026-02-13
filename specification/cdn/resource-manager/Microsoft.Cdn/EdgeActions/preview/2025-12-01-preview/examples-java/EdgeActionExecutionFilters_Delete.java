
/**
 * Samples for EdgeActionExecutionFilters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/EdgeActionExecutionFilters_Delete.json
     */
    /**
     * Sample code: DeleteEdgeActionExecutionFilters.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void
        deleteEdgeActionExecutionFilters(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActionExecutionFilters().delete("testrg", "edgeAction1", "executionFilter1",
            com.azure.core.util.Context.NONE);
    }
}
