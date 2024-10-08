
/**
 * Samples for PrivateEndpointConnection Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: GetPrivateEndpointConnection.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void getPrivateEndpointConnection(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.privateEndpointConnections().getWithResponse("default-azurebatch-japaneast", "sampleacct",
            "testprivateEndpointConnection5testprivateEndpointConnection5.24d6b4b5-e65c-4330-bbe9-3a290d62f8e0",
            com.azure.core.util.Context.NONE);
    }
}
