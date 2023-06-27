import com.azure.core.util.Context;

/** Samples for GatewayRouteConfigs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/GatewayRouteConfigs_List.json
     */
    /**
     * Sample code: GatewayRouteConfigs_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gatewayRouteConfigsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getGatewayRouteConfigs()
            .list("myResourceGroup", "myservice", "default", Context.NONE);
    }
}
