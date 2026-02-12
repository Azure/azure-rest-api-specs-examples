
/**
 * Samples for EdgeActionExecutionFilters Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/EdgeActionExecutionFilters_Get.json
     */
    /**
     * Sample code: GetEdgeActionExecutionFilters.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void getEdgeActionExecutionFilters(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActionExecutionFilters().getWithResponse("testrg", "edgeAction1", "executionFilter1",
            com.azure.core.util.Context.NONE);
    }
}
