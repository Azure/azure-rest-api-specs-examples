import com.azure.core.util.Context;

/** Samples for HybridConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridConnectionGet.json
     */
    /**
     * Sample code: RelayHybridConnectionGet.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayHybridConnectionGet(com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .hybridConnections()
            .getWithResponse("resourcegroup", "example-RelayNamespace-01", "example-Relay-Hybrid-01", Context.NONE);
    }
}
