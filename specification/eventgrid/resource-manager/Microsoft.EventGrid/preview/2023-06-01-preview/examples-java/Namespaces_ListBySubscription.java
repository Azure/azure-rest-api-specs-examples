/** Samples for Namespaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Namespaces_ListBySubscription.json
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
