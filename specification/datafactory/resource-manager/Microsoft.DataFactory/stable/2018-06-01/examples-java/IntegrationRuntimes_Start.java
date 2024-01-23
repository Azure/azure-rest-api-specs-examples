
/**
 * Samples for IntegrationRuntimes Start.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * IntegrationRuntimes_Start.json
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
