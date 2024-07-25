
/**
 * Samples for BillingMeters Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/BillingMeters_Get.json
     */
    /**
     * Sample code: BillingMeters_Get.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void billingMetersGet(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.billingMeters().getWithResponse("East US", com.azure.core.util.Context.NONE);
    }
}
