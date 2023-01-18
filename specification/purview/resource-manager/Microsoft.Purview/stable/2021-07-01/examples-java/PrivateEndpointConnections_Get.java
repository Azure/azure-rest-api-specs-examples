/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/PrivateEndpointConnections_Get.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Get.
     *
     * @param manager Entry point to PurviewManager.
     */
    public static void privateEndpointConnectionsGet(com.azure.resourcemanager.purview.PurviewManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse(
                "SampleResourceGroup", "account1", "privateEndpointConnection1", com.azure.core.util.Context.NONE);
    }
}
