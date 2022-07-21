import com.azure.core.util.Context;

/** Samples for Namespaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/NameSpaces/RelayNameSpaceListBySubscription.json
     */
    /**
     * Sample code: RelayNameSpaceListBySubscription.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayNameSpaceListBySubscription(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.namespaces().list(Context.NONE);
    }
}
