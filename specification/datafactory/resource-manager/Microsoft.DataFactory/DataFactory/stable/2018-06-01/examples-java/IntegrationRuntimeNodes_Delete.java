
/**
 * Samples for IntegrationRuntimeNodes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/IntegrationRuntimeNodes_Delete.json
     */
    /**
     * Sample code: IntegrationRuntimesNodes_Delete.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        integrationRuntimesNodesDelete(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimeNodes().deleteWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleIntegrationRuntime", "Node_1", com.azure.core.util.Context.NONE);
    }
}
