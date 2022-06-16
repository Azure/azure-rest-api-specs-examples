import com.azure.core.util.Context;
import com.azure.resourcemanager.eventgrid.models.ParentType;

/** Samples for PrivateEndpointConnections ListByResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PrivateEndpointConnections_ListByResource.json
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
            .listByResource("examplerg", ParentType.TOPICS, "exampletopic1", null, null, Context.NONE);
    }
}
