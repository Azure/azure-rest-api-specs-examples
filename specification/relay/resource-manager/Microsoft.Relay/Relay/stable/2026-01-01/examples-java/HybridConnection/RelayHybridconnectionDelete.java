
/**
 * Samples for HybridConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/HybridConnection/RelayHybridconnectionDelete.json
     */
    /**
     * Sample code: RelayHybridconnectionDelete.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayHybridconnectionDelete(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.hybridConnections().deleteWithResponse("resourcegroup", "example-RelayNamespace-01",
            "example-Relay-Hybrid-01", com.azure.core.util.Context.NONE);
    }
}
