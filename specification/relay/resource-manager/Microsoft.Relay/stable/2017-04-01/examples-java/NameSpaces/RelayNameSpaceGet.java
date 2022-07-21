import com.azure.core.util.Context;

/** Samples for Namespaces GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/NameSpaces/RelayNameSpaceGet.json
     */
    /**
     * Sample code: RelayNameSpaceGet.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayNameSpaceGet(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.namespaces().getByResourceGroupWithResponse("resourcegroup", "example-RelayNamespace-01", Context.NONE);
    }
}
