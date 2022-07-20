import com.azure.core.util.Context;

/** Samples for WcfRelays Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/Relay/RelayDelete.json
     */
    /**
     * Sample code: RelayDelete.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayDelete(com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .wcfRelays()
            .deleteWithResponse("resourcegroup", "example-RelayNamespace-01", "example-Relay-wcf-01", Context.NONE);
    }
}
