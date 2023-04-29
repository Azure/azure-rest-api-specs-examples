/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/ConfigurationStoresGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_GetConnection.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void privateEndpointConnectionGetConnection(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse("myResourceGroup", "contoso", "myConnection", com.azure.core.util.Context.NONE);
    }
}
