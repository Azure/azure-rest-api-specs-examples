
/**
 * Samples for GatewayRouteConfigs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * GatewayRouteConfigs_Delete.json
     */
    /**
     * Sample code: GatewayRouteConfigs_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gatewayRouteConfigsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getGatewayRouteConfigs().delete("myResourceGroup", "myservice",
            "default", "myRouteConfig", com.azure.core.util.Context.NONE);
    }
}
