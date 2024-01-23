
/**
 * Samples for IntegrationRuntimes Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * IntegrationRuntimes_Stop.json
     */
    /**
     * Sample code: IntegrationRuntimes_Stop.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void integrationRuntimesStop(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimes().stop("exampleResourceGroup", "exampleFactoryName",
            "exampleManagedIntegrationRuntime", com.azure.core.util.Context.NONE);
    }
}
