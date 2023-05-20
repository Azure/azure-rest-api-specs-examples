/** Samples for Namespaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Namespaces_ListByResourceGroup.json
     */
    /**
     * Sample code: Namespaces_ListByResourceGroup.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void namespacesListByResourceGroup(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.namespaces().listByResourceGroup("examplerg", null, null, com.azure.core.util.Context.NONE);
    }
}
