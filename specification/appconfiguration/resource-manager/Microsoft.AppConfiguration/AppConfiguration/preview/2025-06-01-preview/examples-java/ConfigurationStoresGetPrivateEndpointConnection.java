
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_GetConnection.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void privateEndpointConnectionGetConnection(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.privateEndpointConnections().getWithResponse("myResourceGroup", "contoso", "myConnection",
            com.azure.core.util.Context.NONE);
    }
}
