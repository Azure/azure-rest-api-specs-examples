
/**
 * Samples for AppServiceEnvironments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_List.json
     */
    /**
     * Sample code: Get all App Service Environments for a subscription.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getAllAppServiceEnvironmentsForASubscription(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().list(com.azure.core.util.Context.NONE);
    }
}
