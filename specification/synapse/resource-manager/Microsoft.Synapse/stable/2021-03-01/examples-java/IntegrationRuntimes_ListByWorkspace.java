
/**
 * Samples for IntegrationRuntimes ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * IntegrationRuntimes_ListByWorkspace.json
     */
    /**
     * Sample code: List integration runtimes.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void listIntegrationRuntimes(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.integrationRuntimes().listByWorkspace("exampleResourceGroup", "exampleWorkspace",
            com.azure.core.util.Context.NONE);
    }
}
