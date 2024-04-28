
/**
 * Samples for ClientGroups ListByNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * ClientGroups_ListByNamespace.json
     */
    /**
     * Sample code: ClientGroups_ListByNamespace.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void clientGroupsListByNamespace(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.clientGroups().listByNamespace("examplerg", "namespace123", null, null,
            com.azure.core.util.Context.NONE);
    }
}
