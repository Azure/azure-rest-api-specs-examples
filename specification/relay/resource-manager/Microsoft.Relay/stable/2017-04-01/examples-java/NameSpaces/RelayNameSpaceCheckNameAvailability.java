import com.azure.core.util.Context;
import com.azure.resourcemanager.relay.models.CheckNameAvailability;

/** Samples for Namespaces CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/NameSpaces/RelayNameSpaceCheckNameAvailability.json
     */
    /**
     * Sample code: RelayCheckNameAvailability.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayCheckNameAvailability(com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .namespaces()
            .checkNameAvailabilityWithResponse(new CheckNameAvailability().withName("sdk-Namespace1321"), Context.NONE);
    }
}
