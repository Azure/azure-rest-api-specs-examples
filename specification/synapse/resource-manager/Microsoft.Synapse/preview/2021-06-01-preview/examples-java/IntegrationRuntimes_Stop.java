import com.azure.core.util.Context;

/** Samples for IntegrationRuntimes Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Stop.json
     */
    /**
     * Sample code: Stop integration runtime.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void stopIntegrationRuntime(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimes()
            .stop("exampleResourceGroup", "exampleWorkspace", "exampleManagedIntegrationRuntime", Context.NONE);
    }
}
