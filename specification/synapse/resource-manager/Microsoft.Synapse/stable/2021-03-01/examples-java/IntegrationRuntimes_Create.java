
import com.azure.resourcemanager.synapse.models.SelfHostedIntegrationRuntime;

/**
 * Samples for IntegrationRuntimes Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/IntegrationRuntimes_Create.
     * json
     */
    /**
     * Sample code: Create integration runtime.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void createIntegrationRuntime(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.integrationRuntimes().define("exampleIntegrationRuntime")
            .withExistingWorkspace("exampleResourceGroup", "exampleWorkspace")
            .withProperties(new SelfHostedIntegrationRuntime().withDescription("A selfhosted integration runtime"))
            .create();
    }
}
