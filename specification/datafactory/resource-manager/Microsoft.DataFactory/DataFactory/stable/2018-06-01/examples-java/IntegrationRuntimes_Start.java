
/**
 * Samples for IntegrationRuntimes Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/IntegrationRuntimes_Start.json
     */
    /**
     * Sample code: IntegrationRuntimes_Start.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void integrationRuntimesStart(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimes().start("exampleResourceGroup", "exampleFactoryName",
            "exampleManagedIntegrationRuntime", com.azure.core.util.Context.NONE);
    }
}
