
/**
 * Samples for IntegrationRuntimeAuthKeysOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * IntegrationRuntimes_ListAuthKeys.json
     */
    /**
     * Sample code: List auth keys.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void listAuthKeys(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.integrationRuntimeAuthKeysOperations().listWithResponse("exampleResourceGroup", "exampleWorkspace",
            "exampleIntegrationRuntime", com.azure.core.util.Context.NONE);
    }
}
