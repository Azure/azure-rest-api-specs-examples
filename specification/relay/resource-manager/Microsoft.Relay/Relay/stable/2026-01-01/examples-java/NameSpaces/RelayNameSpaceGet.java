
/**
 * Samples for Namespaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NameSpaces/RelayNameSpaceGet.json
     */
    /**
     * Sample code: RelayNameSpaceGet.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayNameSpaceGet(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.namespaces().getByResourceGroupWithResponse("RG-eg", "example-RelayRelayNamespace-01",
            com.azure.core.util.Context.NONE);
    }
}
