
/**
 * Samples for AppResiliency Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/AppResiliency_Get.json
     */
    /**
     * Sample code: Get App Resiliency.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getAppResiliency(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.appResiliencies().getWithResponse("rg", "testcontainerApp0", "resiliency-policy-1",
            com.azure.core.util.Context.NONE);
    }
}
