
import com.azure.resourcemanager.relay.models.Relaytype;

/**
 * Samples for WcfRelays CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/Relay/RelayCreate.json
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
