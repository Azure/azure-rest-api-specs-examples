import com.azure.resourcemanager.eventgrid.models.PrivateEndpointConnectionsParentType;

/** Samples for PrivateEndpointConnections ListByResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PrivateEndpointConnections_ListByResource.json
     */
    /**
     * Sample code: PrivateEndpointConnections_ListByResource.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void privateEndpointConnectionsListByResource(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .privateEndpointConnections()
            .listByResource(
                "examplerg",
                PrivateEndpointConnectionsParentType.TOPICS,
                "exampletopic1",
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
