
import com.azure.resourcemanager.relay.models.Relaytype;

/**
 * Samples for WcfRelays CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/Relay/RelayCreate.json
     */
    /**
     * Sample code: RelayCreate.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayCreate(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.wcfRelays().define("example-Relay-Wcf-1194")
            .withExistingNamespace("resourcegroup", "example-RelayNamespace-9953").withRelayType(Relaytype.NET_TCP)
            .withRequiresClientAuthorization(true).withRequiresTransportSecurity(true).create();
    }
}
