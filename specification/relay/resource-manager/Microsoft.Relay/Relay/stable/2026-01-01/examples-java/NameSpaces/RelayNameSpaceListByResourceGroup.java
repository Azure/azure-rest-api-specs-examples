
/**
 * Samples for Namespaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NameSpaces/RelayNameSpaceListByResourceGroup.json
     */
    /**
     * Sample code: RelayNameSpaceListByResourceGroup.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayNameSpaceListByResourceGroup(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.namespaces().listByResourceGroup("resourcegroup", com.azure.core.util.Context.NONE);
    }
}
