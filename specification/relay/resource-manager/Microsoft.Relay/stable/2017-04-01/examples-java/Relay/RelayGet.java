import com.azure.core.util.Context;

/** Samples for WcfRelays Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/Relay/RelayGet.json
     */
    /**
     * Sample code: RelayGet.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayGet(com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .wcfRelays()
            .getWithResponse("resourcegroup", "example-RelayNamespace-9953", "example-Relay-Wcf-1194", Context.NONE);
    }
}
