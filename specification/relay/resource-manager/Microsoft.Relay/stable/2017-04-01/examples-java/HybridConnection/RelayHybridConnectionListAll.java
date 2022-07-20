import com.azure.core.util.Context;

/** Samples for HybridConnections ListByNamespace. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridConnectionListAll.json
     */
    /**
     * Sample code: RelayHybridConnectionListAll.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayHybridConnectionListAll(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.hybridConnections().listByNamespace("resourcegroup", "example-RelayNamespace-01", Context.NONE);
    }
}
