
import com.azure.resourcemanager.eventgrid.models.PrivateEndpointConnectionsParentType;

/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * PrivateEndpointConnections_Delete.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Delete.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void privateEndpointConnectionsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.privateEndpointConnections().delete("examplerg", PrivateEndpointConnectionsParentType.TOPICS,
            "exampletopic1", "BMTPE5.8A30D251-4C61-489D-A1AA-B37C4A329B8B", com.azure.core.util.Context.NONE);
    }
}
