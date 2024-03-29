
/**
 * Samples for PrivateEndpointConnection ListByBatchAccount.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/PrivateEndpointConnectionsList.
     * json
     */
    /**
     * Sample code: ListPrivateEndpointConnections.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void listPrivateEndpointConnections(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.privateEndpointConnections().listByBatchAccount("default-azurebatch-japaneast", "sampleacct", null,
            com.azure.core.util.Context.NONE);
    }
}
