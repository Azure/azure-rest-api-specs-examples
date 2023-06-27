import com.azure.core.util.Context;

/** Samples for GatewayRouteConfigs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/GatewayRouteConfigs_Get.json
     */
    /**
     * Sample code: GatewayRouteConfigs_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gatewayRouteConfigsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getGatewayRouteConfigs()
            .getWithResponse("myResourceGroup", "myservice", "default", "myRouteConfig", Context.NONE);
    }
}
