
/**
 * Samples for AppServicePlans GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetAppServicePlan.json
     */
    /**
     * Sample code: Get App Service plan.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAppServicePlan(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServicePlans().getByResourceGroupWithResponse("testrg123", "testsf6141",
            com.azure.core.util.Context.NONE);
    }
}
