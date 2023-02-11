/** Samples for IntegrationRuntimeStatusOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_GetStatus.json
     */
    /**
     * Sample code: Get status.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getStatus(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimeStatusOperations()
            .getWithResponse(
                "exampleResourceGroup",
                "exampleWorkspace",
                "exampleIntegrationRuntime",
                com.azure.core.util.Context.NONE);
    }
}
