
/**
 * Samples for IntegrationRuntimes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/
     * IntegrationRuntimes_Get.json
     */
    /**
     * Sample code: Get integration runtime.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void getIntegrationRuntime(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.integrationRuntimes().getWithResponse("exampleResourceGroup", "exampleWorkspace",
            "exampleIntegrationRuntime", null, com.azure.core.util.Context.NONE);
    }
}
