
/**
 * Samples for DaprSubscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * DaprSubscriptions_Get_BulkSubscribeAndScopes.json
     */
    /**
     * Sample code: Get Dapr subscription with default route only.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getDaprSubscriptionWithDefaultRouteOnly(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprSubscriptions().getWithResponse("examplerg", "myenvironment", "mypubsubcomponent",
            com.azure.core.util.Context.NONE);
    }
}
