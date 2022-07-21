import com.azure.core.util.Context;

/** Samples for HybridConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridconnectionDelete.json
     */
    /**
     * Sample code: RelayHybridconnectionDelete.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayHybridconnectionDelete(com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .hybridConnections()
            .deleteWithResponse("resourcegroup", "example-RelayNamespace-01", "example-Relay-Hybrid-01", Context.NONE);
    }
}
