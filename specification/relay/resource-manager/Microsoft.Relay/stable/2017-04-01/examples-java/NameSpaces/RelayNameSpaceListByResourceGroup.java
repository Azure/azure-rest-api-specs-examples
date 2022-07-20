import com.azure.core.util.Context;

/** Samples for Namespaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/NameSpaces/RelayNameSpaceListByResourceGroup.json
     */
    /**
     * Sample code: RelayNameSpaceListByResourceGroup.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayNameSpaceListByResourceGroup(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.namespaces().listByResourceGroup("resourcegroup", Context.NONE);
    }
}
