
/**
 * Samples for IntegrationRuntimeStatusOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/IntegrationRuntimes_GetStatus
     * .json
     */
    /**
     * Sample code: Get status.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void getStatus(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.integrationRuntimeStatusOperations().getWithResponse("exampleResourceGroup", "exampleWorkspace",
            "exampleIntegrationRuntime", com.azure.core.util.Context.NONE);
    }
}
