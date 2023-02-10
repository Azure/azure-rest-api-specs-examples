/** Samples for KustoOperations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoOperationsList.json
     */
    /**
     * Sample code: KustoOperationsList.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoOperationsList(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.kustoOperations().list(com.azure.core.util.Context.NONE);
    }
}
