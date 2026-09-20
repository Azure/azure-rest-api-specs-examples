
/**
 * Samples for HttpRouteConfig Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/HttpRouteConfig_Get.json
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
