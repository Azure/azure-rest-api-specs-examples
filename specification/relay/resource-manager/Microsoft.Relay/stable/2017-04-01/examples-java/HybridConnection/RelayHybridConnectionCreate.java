/** Samples for HybridConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridConnectionCreate.json
     */
    /**
     * Sample code: RelayHybridConnectionCreate.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayHybridConnectionCreate(com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .hybridConnections()
            .define("example-Relay-Hybrid-01")
            .withExistingNamespace("resourcegroup", "example-RelayNamespace-01")
            .withRequiresClientAuthorization(true)
            .create();
    }
}
