
/**
 * Samples for PrivateLinkResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/PrivateEndpointConnections/PrivateLinkResourcesList.json
     */
    /**
     * Sample code: NameSpacePrivateLinkResourcesGet.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void nameSpacePrivateLinkResourcesGet(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.privateLinkResources().listWithResponse("resourcegroup", "example-RelayNamespace-5849",
            com.azure.core.util.Context.NONE);
    }
}
