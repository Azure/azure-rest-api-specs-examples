
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/PrivateEndpointConnections/PrivateEndpointConnectionsGet.json
     */
    /**
     * Sample code: NameSpacePrivateEndPointConnectionGet.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void nameSpacePrivateEndPointConnectionGet(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.privateEndpointConnections().getWithResponse("myResourceGroup", "example-RelayNamespace-5849",
            "{privateEndpointConnection name}", com.azure.core.util.Context.NONE);
    }
}
