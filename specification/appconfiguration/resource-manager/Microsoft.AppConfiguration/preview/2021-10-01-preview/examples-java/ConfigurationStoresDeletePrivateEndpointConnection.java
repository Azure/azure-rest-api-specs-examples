import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/ConfigurationStoresDeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Delete.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void privateEndpointConnectionsDelete(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.privateEndpointConnections().delete("myResourceGroup", "contoso", "myConnection", Context.NONE);
    }
}
