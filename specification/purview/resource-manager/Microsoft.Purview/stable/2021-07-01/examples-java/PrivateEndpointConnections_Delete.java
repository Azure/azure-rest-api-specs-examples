/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/PrivateEndpointConnections_Delete.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Delete.
     *
     * @param manager Entry point to PurviewManager.
     */
    public static void privateEndpointConnectionsDelete(com.azure.resourcemanager.purview.PurviewManager manager) {
        manager
            .privateEndpointConnections()
            .delete("SampleResourceGroup", "account1", "privateEndpointConnection1", com.azure.core.util.Context.NONE);
    }
}
