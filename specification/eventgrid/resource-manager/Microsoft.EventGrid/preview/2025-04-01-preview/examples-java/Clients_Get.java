
/**
 * Samples for Clients Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Clients_Get.json
     */
    /**
     * Sample code: Clients_Get.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void clientsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.clients().getWithResponse("examplerg", "exampleNamespaceName1", "exampleClientName1",
            com.azure.core.util.Context.NONE);
    }
}
