
/**
 * Samples for ClientGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/ClientGroups_Delete.json
     */
    /**
     * Sample code: ClientGroups_Delete.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void clientGroupsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.clientGroups().delete("examplerg", "exampleNamespaceName1", "exampleClientGroupName1",
            com.azure.core.util.Context.NONE);
    }
}
