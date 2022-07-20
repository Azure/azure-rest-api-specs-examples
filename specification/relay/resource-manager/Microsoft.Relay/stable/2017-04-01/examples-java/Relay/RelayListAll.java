import com.azure.core.util.Context;

/** Samples for WcfRelays ListByNamespace. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/Relay/RelayListAll.json
     */
    /**
     * Sample code: RelayListAll.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayListAll(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.wcfRelays().listByNamespace("resourcegroup", "example-RelayNamespace-01", Context.NONE);
    }
}
