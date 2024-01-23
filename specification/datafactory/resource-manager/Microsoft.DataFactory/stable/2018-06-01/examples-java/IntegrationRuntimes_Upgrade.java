
/**
 * Samples for IntegrationRuntimes Upgrade.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * IntegrationRuntimes_Upgrade.json
     */
    /**
     * Sample code: IntegrationRuntimes_Upgrade.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void integrationRuntimesUpgrade(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimes().upgradeWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleIntegrationRuntime", com.azure.core.util.Context.NONE);
    }
}
