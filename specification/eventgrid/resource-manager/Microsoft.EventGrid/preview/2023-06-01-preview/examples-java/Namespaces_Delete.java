/** Samples for Namespaces Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Namespaces_Delete.json
     */
    /**
     * Sample code: Namespaces_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void namespacesDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.namespaces().delete("examplerg", "exampleNamespaceName1", com.azure.core.util.Context.NONE);
    }
}
