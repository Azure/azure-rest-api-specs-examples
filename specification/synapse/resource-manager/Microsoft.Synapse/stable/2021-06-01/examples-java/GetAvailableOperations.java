/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetAvailableOperations.json
     */
    /**
     * Sample code: Get available operations.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getAvailableOperations(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.operations().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
