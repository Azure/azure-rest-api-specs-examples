
import com.azure.resourcemanager.edgeactions.models.EdgeActionExecutionFilterProperties;

/**
 * Samples for EdgeActionExecutionFilters Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeActionExecutionFilters_Create.json
     */
    /**
     * Sample code: CreateEdgeActionExecutionFilters.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void
        createEdgeActionExecutionFilters(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActionExecutionFilters().define("executionFilter1").withRegion("global")
            .withExistingEdgeAction("testrg", "edgeAction1")
            .withProperties(new EdgeActionExecutionFilterProperties().withVersionId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Cdn/edgeActions/edgeAction1/versions/version1")
                .withExecutionFilterIdentifierHeaderName("header-key")
                .withExecutionFilterIdentifierHeaderValue("header-value"))
            .create();
    }
}
