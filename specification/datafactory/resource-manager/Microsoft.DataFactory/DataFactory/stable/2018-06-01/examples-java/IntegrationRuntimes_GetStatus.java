
/**
 * Samples for IntegrationRuntimes GetStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/IntegrationRuntimes_GetStatus.json
     */
    /**
     * Sample code: IntegrationRuntimes_GetStatus.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void integrationRuntimesGetStatus(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimes().getStatusWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleIntegrationRuntime", com.azure.core.util.Context.NONE);
    }
}
