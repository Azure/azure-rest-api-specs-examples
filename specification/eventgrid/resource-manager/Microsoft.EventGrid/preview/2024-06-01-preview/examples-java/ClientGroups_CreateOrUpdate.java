
/**
 * Samples for ClientGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * ClientGroups_CreateOrUpdate.json
     */
    /**
     * Sample code: ClientGroups_CreateOrUpdate.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void clientGroupsCreateOrUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.clientGroups().define("exampleClientGroupName1")
            .withExistingNamespace("examplerg", "exampleNamespaceName1").withDescription("This is a test client group")
            .withQuery("attributes.b IN ['a', 'b', 'c']").create();
    }
}
