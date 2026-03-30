
/**
 * Samples for AppServicePlans Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeleteAppServicePlan.json
     */
    /**
     * Sample code: Delete App Service plan.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void deleteAppServicePlan(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServicePlans().deleteWithResponse("testrg123", "testsf6141",
            com.azure.core.util.Context.NONE);
    }
}
