/** Samples for IntegrationRuntimes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Delete.json
     */
    /**
     * Sample code: Delete integration runtime.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void deleteIntegrationRuntime(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimes()
            .delete(
                "exampleResourceGroup",
                "exampleWorkspace",
                "exampleIntegrationRuntime",
                com.azure.core.util.Context.NONE);
    }
}
