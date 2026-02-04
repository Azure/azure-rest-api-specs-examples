
/**
 * Samples for PrivateEndpointConnections ListByConfigurationStore.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresListPrivateEndpointConnections.json
     */
    /**
     * Sample code: PrivateEndpointConnection_List.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void
        privateEndpointConnectionList(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.privateEndpointConnections().listByConfigurationStore("myResourceGroup", "contoso",
            com.azure.core.util.Context.NONE);
    }
}
