
/**
 * Samples for GetUsagesInLocation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetUsagesInLocation.json
     */
    /**
     * Sample code: Get usages in location for subscription.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getUsagesInLocationForSubscription(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getGetUsagesInLocations().list("West US", com.azure.core.util.Context.NONE);
    }
}
