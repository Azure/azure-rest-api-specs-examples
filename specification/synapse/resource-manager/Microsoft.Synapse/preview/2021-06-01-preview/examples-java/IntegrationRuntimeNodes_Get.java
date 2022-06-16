import com.azure.core.util.Context;

/** Samples for IntegrationRuntimeNodes Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimeNodes_Get.json
     */
    /**
     * Sample code: Get integration runtime node.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getIntegrationRuntimeNode(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimeNodes()
            .getWithResponse(
                "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", "Node_1", Context.NONE);
    }
}
