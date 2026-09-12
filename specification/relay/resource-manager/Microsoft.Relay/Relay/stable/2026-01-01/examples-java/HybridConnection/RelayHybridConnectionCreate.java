
/**
 * Samples for HybridConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/HybridConnection/RelayHybridConnectionCreate.json
     */
    /**
     * Sample code: RelayHybridConnectionCreate.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayHybridConnectionCreate(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.hybridConnections().define("example-Relay-Hybrid-01")
            .withExistingNamespace("resourcegroup", "example-RelayNamespace-01").withRequiresClientAuthorization(true)
            .create();
    }
}
