
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresDeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Delete.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void
        privateEndpointConnectionsDelete(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.privateEndpointConnections().delete("myResourceGroup", "contoso", "myConnection",
            com.azure.core.util.Context.NONE);
    }
}
