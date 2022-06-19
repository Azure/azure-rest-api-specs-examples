import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections ListByConfigurationStore. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresListPrivateEndpointConnections.json
     */
    /**
     * Sample code: PrivateEndpointConnection_List.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void privateEndpointConnectionList(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.privateEndpointConnections().listByConfigurationStore("myResourceGroup", "contoso", Context.NONE);
    }
}
