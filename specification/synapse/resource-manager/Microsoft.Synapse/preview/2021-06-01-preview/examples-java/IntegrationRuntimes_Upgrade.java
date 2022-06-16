import com.azure.core.util.Context;

/** Samples for IntegrationRuntimes Upgrade. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Upgrade.json
     */
    /**
     * Sample code: Upgrade integration runtime.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void upgradeIntegrationRuntime(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimes()
            .upgradeWithResponse("exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", Context.NONE);
    }
}
