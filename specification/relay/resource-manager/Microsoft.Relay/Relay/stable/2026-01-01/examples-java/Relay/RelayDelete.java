
/**
 * Samples for WcfRelays Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/Relay/RelayDelete.json
     */
    /**
     * Sample code: RelayDelete.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayDelete(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.wcfRelays().deleteWithResponse("resourcegroup", "example-RelayNamespace-01", "example-Relay-wcf-01",
            com.azure.core.util.Context.NONE);
    }
}
