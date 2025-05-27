
/**
 * Samples for IntegrationRuntimes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * IntegrationRuntimes_Delete.json
     */
    /**
     * Sample code: IntegrationRuntimes_Delete.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void integrationRuntimesDelete(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimes().deleteWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleIntegrationRuntime", com.azure.core.util.Context.NONE);
    }
}
