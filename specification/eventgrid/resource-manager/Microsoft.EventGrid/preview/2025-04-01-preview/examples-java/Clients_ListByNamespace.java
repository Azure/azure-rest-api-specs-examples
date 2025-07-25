
/**
 * Samples for Clients ListByNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * Clients_ListByNamespace.json
     */
    /**
     * Sample code: Clients_ListByNamespace.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void clientsListByNamespace(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.clients().listByNamespace("examplerg", "namespace123", null, null, com.azure.core.util.Context.NONE);
    }
}
