
import com.azure.resourcemanager.datafactory.models.UpdateIntegrationRuntimeNodeRequest;

/**
 * Samples for IntegrationRuntimeNodes UpdateSync.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * IntegrationRuntimeNodes_Update.json
     */
    /**
     * Sample code: IntegrationRuntimeNodes_Update.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void integrationRuntimeNodesUpdate(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimeNodes().updateWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleIntegrationRuntime", "Node_1", new UpdateIntegrationRuntimeNodeRequest().withConcurrentJobsLimit(2),
            com.azure.core.util.Context.NONE);
    }
}
