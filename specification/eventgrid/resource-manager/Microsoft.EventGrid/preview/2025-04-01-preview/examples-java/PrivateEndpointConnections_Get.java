
import com.azure.resourcemanager.eventgrid.models.PrivateEndpointConnectionsParentType;

/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * PrivateEndpointConnections_Get.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Get.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void privateEndpointConnectionsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.privateEndpointConnections().getWithResponse("examplerg", PrivateEndpointConnectionsParentType.TOPICS,
            "exampletopic1", "BMTPE5.8A30D251-4C61-489D-A1AA-B37C4A329B8B", com.azure.core.util.Context.NONE);
    }
}
