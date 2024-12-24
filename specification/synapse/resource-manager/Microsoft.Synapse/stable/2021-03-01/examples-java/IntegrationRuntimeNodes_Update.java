
import com.azure.resourcemanager.synapse.models.UpdateIntegrationRuntimeNodeRequest;

/**
 * Samples for IntegrationRuntimeNodes Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * IntegrationRuntimeNodes_Update.json
     */
    /**
     * Sample code: Update integration runtime node.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void updateIntegrationRuntimeNode(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.integrationRuntimeNodes().updateWithResponse("exampleResourceGroup", "exampleWorkspace",
            "exampleIntegrationRuntime", "Node_1", new UpdateIntegrationRuntimeNodeRequest().withConcurrentJobsLimit(2),
            com.azure.core.util.Context.NONE);
    }
}
