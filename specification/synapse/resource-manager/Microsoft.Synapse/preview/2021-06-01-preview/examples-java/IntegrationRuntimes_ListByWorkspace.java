/** Samples for IntegrationRuntimes ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_ListByWorkspace.json
     */
    /**
     * Sample code: List integration runtimes.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listIntegrationRuntimes(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimes()
            .listByWorkspace("exampleResourceGroup", "exampleWorkspace", com.azure.core.util.Context.NONE);
    }
}
