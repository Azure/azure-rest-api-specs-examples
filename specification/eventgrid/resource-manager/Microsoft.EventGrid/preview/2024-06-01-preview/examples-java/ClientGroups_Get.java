
/**
 * Samples for ClientGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/ClientGroups_Get
     * .json
     */
    /**
     * Sample code: ClientGroups_Get.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void clientGroupsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.clientGroups().getWithResponse("examplerg", "exampleNamespaceName1", "exampleClientGroupName1",
            com.azure.core.util.Context.NONE);
    }
}
