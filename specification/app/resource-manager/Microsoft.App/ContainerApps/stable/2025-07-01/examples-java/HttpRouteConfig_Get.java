
/**
 * Samples for HttpRouteConfig Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/HttpRouteConfig_Get.
     * json
     */
    /**
     * Sample code: Get HttpRoute.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getHttpRoute(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.httpRouteConfigs().getWithResponse("examplerg", "testcontainerenv", "httproutefriendlyname",
            com.azure.core.util.Context.NONE);
    }
}
