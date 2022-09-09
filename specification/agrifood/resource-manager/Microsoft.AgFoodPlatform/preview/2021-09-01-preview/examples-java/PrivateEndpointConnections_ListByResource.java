import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections ListByResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/PrivateEndpointConnections_ListByResource.json
     */
    /**
     * Sample code: PrivateEndpointConnections_ListByResource.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void privateEndpointConnectionsListByResource(
        com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager
            .privateEndpointConnections()
            .listByResource("examples-rg", "examples-farmbeatsResourceName", Context.NONE);
    }
}
