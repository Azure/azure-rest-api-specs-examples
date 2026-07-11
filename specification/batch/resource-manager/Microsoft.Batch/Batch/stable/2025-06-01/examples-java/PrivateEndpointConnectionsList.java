
/**
 * Samples for PrivateEndpointConnection ListByBatchAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/PrivateEndpointConnectionsList.json
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
