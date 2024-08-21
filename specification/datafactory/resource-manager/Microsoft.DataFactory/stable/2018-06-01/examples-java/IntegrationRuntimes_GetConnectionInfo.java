
/**
 * Samples for IntegrationRuntimes GetConnectionInfo.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * IntegrationRuntimes_GetConnectionInfo.json
     */
    /**
     * Sample code: IntegrationRuntimes_GetConnectionInfo.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        integrationRuntimesGetConnectionInfo(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimes().getConnectionInfoWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleIntegrationRuntime", com.azure.core.util.Context.NONE);
    }
}
