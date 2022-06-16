import com.azure.core.util.Context;

/** Samples for IntegrationRuntimes Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Start.json
     */
    /**
     * Sample code: Start integration runtime.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void startIntegrationRuntime(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimes()
            .start("exampleResourceGroup", "exampleWorkspace", "exampleManagedIntegrationRuntime", Context.NONE);
    }
}
