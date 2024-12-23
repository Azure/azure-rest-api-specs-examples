
/**
 * Samples for IntegrationRuntimes EnableInteractiveQuery.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * IntegrationRuntimes_EnableInteractiveQuery.json
     */
    /**
     * Sample code: Stop integration runtime.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void stopIntegrationRuntime(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.integrationRuntimes().enableInteractiveQuery("exampleResourceGroup", "exampleWorkspace",
            "exampleManagedIntegrationRuntime", com.azure.core.util.Context.NONE);
    }
}
