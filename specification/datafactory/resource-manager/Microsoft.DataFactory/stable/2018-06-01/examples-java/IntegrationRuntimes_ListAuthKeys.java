
/**
 * Samples for IntegrationRuntimes ListAuthKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * IntegrationRuntimes_ListAuthKeys.json
     */
    /**
     * Sample code: IntegrationRuntimes_ListAuthKeys.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        integrationRuntimesListAuthKeys(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimes().listAuthKeysWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleIntegrationRuntime", com.azure.core.util.Context.NONE);
    }
}
