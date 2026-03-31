
/**
 * Samples for Provider GetFunctionAppStacksForLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetFunctionAppStacksForLocation.json
     */
    /**
     * Sample code: Get Locations Function App Stacks.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getLocationsFunctionAppStacks(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getProviders().getFunctionAppStacksForLocation("westus", null,
            com.azure.core.util.Context.NONE);
    }
}
