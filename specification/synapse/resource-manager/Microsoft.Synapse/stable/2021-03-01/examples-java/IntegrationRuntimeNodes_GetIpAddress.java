
/**
 * Samples for IntegrationRuntimeNodeIpAddressOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * IntegrationRuntimeNodes_GetIpAddress.json
     */
    /**
     * Sample code: Get integration runtime node IP address.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void getIntegrationRuntimeNodeIPAddress(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.integrationRuntimeNodeIpAddressOperations().getWithResponse("exampleResourceGroup", "exampleWorkspace",
            "exampleIntegrationRuntime", "Node_1", com.azure.core.util.Context.NONE);
    }
}
