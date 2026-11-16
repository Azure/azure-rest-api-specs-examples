
/**
 * Samples for Namespaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/EHNameSpaceList.json
     */
    /**
     * Sample code: NamespacesListBySubscription.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void namespacesListBySubscription(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNamespaces().list(com.azure.core.util.Context.NONE);
    }
}
