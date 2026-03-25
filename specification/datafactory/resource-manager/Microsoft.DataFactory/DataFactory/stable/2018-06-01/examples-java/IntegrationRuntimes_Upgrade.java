
/**
 * Samples for IntegrationRuntimes Upgrade.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/IntegrationRuntimes_Upgrade.json
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
