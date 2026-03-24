
/**
 * Samples for IntegrationRuntimeOperation DisableInteractiveQuery.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/IntegrationRuntimes_DisableInteractiveQuery.json
     */
    /**
     * Sample code: IntegrationRuntime_DisableInteractiveQuery.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        integrationRuntimeDisableInteractiveQuery(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimeOperations().disableInteractiveQuery("exampleResourceGroup", "exampleFactoryName",
            "exampleIntegrationRuntime", com.azure.core.util.Context.NONE);
    }
}
