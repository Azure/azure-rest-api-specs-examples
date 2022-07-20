import com.azure.core.util.Context;

/** Samples for Namespaces Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/NameSpaces/RelayNameSpaceDelete.json
     */
    /**
     * Sample code: RelayNameSpaceDelete.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayNameSpaceDelete(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.namespaces().delete("resourcegroup", "example-RelayNamespace-01", Context.NONE);
    }
}
