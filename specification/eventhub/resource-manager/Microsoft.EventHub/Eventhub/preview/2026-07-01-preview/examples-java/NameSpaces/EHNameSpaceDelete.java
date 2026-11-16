
/**
 * Samples for Namespaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/EHNameSpaceDelete.json
     */
    /**
     * Sample code: NameSpaceDelete.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void nameSpaceDelete(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNamespaces().delete("ResurceGroupSample", "NamespaceSample",
            com.azure.core.util.Context.NONE);
    }
}
