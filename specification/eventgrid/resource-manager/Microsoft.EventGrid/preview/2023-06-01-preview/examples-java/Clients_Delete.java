/** Samples for Clients Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Clients_Delete.json
     */
    /**
     * Sample code: Clients_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void clientsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .clients()
            .delete("examplerg", "exampleNamespaceName1", "exampleClientName1", com.azure.core.util.Context.NONE);
    }
}
