
/**
 * Samples for Namespaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/EHNameSpaceListByResourceGroup.json
     */
    /**
     * Sample code: NamespaceListByResourceGroup.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void namespaceListByResourceGroup(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNamespaces().listByResourceGroup("ResurceGroupSample",
            com.azure.core.util.Context.NONE);
    }
}
