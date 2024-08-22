
/**
 * Samples for IntegrationRuntimes GetMonitoringData.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * IntegrationRuntimes_GetMonitoringData.json
     */
    /**
     * Sample code: IntegrationRuntimes_GetMonitoringData.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        integrationRuntimesGetMonitoringData(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimes().getMonitoringDataWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleIntegrationRuntime", com.azure.core.util.Context.NONE);
    }
}
