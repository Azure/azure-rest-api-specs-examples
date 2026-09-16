
/**
 * Samples for DataConnectorDefinitions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectorDefinitions/DeleteDataConnectorDefinitionById.json
     */
    /**
     * Sample code: Delete data connector definition.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        deleteDataConnectorDefinition(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectorDefinitions().deleteWithResponse("myRg", "myWorkspace",
            "73e01a99-5cd7-4139-a149-9f2736ff2ab5", com.azure.core.util.Context.NONE);
    }
}
