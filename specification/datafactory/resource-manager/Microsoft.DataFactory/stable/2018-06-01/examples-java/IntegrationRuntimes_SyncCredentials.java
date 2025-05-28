
/**
 * Samples for IntegrationRuntimes SyncCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * IntegrationRuntimes_SyncCredentials.json
     */
    /**
     * Sample code: IntegrationRuntimes_SyncCredentials.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        integrationRuntimesSyncCredentials(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimes().syncCredentialsWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleIntegrationRuntime", com.azure.core.util.Context.NONE);
    }
}
