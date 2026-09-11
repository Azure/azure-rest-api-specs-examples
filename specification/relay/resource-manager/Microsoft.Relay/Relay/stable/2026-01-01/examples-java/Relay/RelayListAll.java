
/**
 * Samples for WcfRelays ListByNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/Relay/RelayListAll.json
     */
    /**
     * Sample code: RelayListAll.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayListAll(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.wcfRelays().listByNamespace("resourcegroup", "example-RelayNamespace-01",
            com.azure.core.util.Context.NONE);
    }
}
