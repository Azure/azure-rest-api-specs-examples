
/**
 * Samples for IntegrationRuntimes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * IntegrationRuntimes_Get.json
     */
    /**
     * Sample code: IntegrationRuntimes_Get.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void integrationRuntimesGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimes().getWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleIntegrationRuntime", null, com.azure.core.util.Context.NONE);
    }
}
