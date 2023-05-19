/** Samples for PrivateLinkResources ListByResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PrivateLinkResources_ListByResource.json
     */
    /**
     * Sample code: PrivateLinkResources_ListByResource.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void privateLinkResourcesListByResource(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .privateLinkResources()
            .listByResource("examplerg", "topics", "exampletopic1", null, null, com.azure.core.util.Context.NONE);
    }
}
