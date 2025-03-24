
/**
 * Samples for Namespaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * Namespaces_ListBySubscription.json
     */
    /**
     * Sample code: Namespaces_ListBySubscription.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void namespacesListBySubscription(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.namespaces().list(null, null, com.azure.core.util.Context.NONE);
    }
}
