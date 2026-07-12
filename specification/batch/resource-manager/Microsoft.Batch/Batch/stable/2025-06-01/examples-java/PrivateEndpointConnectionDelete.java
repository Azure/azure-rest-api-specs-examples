
/**
 * Samples for PrivateEndpointConnection Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/PrivateEndpointConnectionDelete.json
     */
    /**
     * Sample code: PrivateEndpointConnectionDelete.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void privateEndpointConnectionDelete(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.privateEndpointConnections().delete("default-azurebatch-japaneast", "sampleacct",
            "testprivateEndpointConnection5testprivateEndpointConnection5.24d6b4b5-e65c-4330-bbe9-3a290d62f8e0",
            com.azure.core.util.Context.NONE);
    }
}
