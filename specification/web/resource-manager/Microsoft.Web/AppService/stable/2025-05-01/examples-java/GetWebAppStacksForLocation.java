
/**
 * Samples for Provider GetWebAppStacksForLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetWebAppStacksForLocation.json
     */
    /**
     * Sample code: Get Locations Web App Stacks.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getLocationsWebAppStacks(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getProviders().getWebAppStacksForLocation("westus", null,
            com.azure.core.util.Context.NONE);
    }
}
