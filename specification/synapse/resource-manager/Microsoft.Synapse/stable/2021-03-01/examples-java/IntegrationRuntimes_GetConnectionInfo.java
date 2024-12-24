
/**
 * Samples for IntegrationRuntimeConnectionInfos Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * IntegrationRuntimes_GetConnectionInfo.json
     */
    /**
     * Sample code: Get connection info.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void getConnectionInfo(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.integrationRuntimeConnectionInfos().getWithResponse("exampleResourceGroup", "exampleWorkspace",
            "exampleIntegrationRuntime", com.azure.core.util.Context.NONE);
    }
}
