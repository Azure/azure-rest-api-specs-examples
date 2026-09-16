
/**
 * Samples for DataConnectorDefinitions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectorDefinitions/GetCustomizableDataConnectorDefinitionById.json
     */
    /**
     * Sample code: Get customize data connector definition.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getCustomizeDataConnectorDefinition(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectorDefinitions().getWithResponse("myRg", "myWorkspace",
            "763f9fa1-c2d3-4fa2-93e9-bccd4899aa12", com.azure.core.util.Context.NONE);
    }
}
