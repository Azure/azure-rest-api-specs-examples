
/**
 * Samples for HybridConnections ListByNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/HybridConnection/RelayHybridConnectionListAll.json
     */
    /**
     * Sample code: RelayHybridConnectionListAll.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayHybridConnectionListAll(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.hybridConnections().listByNamespace("resourcegroup", "example-RelayNamespace-01",
            com.azure.core.util.Context.NONE);
    }
}
