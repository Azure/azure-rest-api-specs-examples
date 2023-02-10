/** Samples for IntegrationRuntimeCredentials Sync. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_SyncCredentials.json
     */
    /**
     * Sample code: Sync credentials.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void syncCredentials(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimeCredentials()
            .syncWithResponse(
                "exampleResourceGroup",
                "exampleWorkspace",
                "exampleIntegrationRuntime",
                com.azure.core.util.Context.NONE);
    }
}
