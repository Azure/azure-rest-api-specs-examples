
/**
 * Samples for AppServicePlans List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListAppServicePlans.json
     */
    /**
     * Sample code: List App Service plans.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAppServicePlans(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServicePlans().list(null, com.azure.core.util.Context.NONE);
    }
}
