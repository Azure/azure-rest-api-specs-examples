
import com.azure.resourcemanager.edgeactions.models.EdgeActionExecutionFilter;
import com.azure.resourcemanager.edgeactions.models.EdgeActionExecutionFilterUpdateProperties;

/**
 * Samples for EdgeActionExecutionFilters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/EdgeActionExecutionFilters_Update.json
     */
    /**
     * Sample code: UpdateEdgeActionExecutionFilters.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void
        updateEdgeActionExecutionFilters(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        EdgeActionExecutionFilter resource = manager.edgeActionExecutionFilters()
            .getWithResponse("testrg", "edgeAction1", "executionFilter1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(
            new EdgeActionExecutionFilterUpdateProperties().withExecutionFilterIdentifierHeaderValue("header-value2"))
            .apply();
    }
}
