
/**
 * Samples for Provider GetWebAppStacks.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetWebAppStacks.json
     */
    /**
     * Sample code: Get Web App Stacks.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getWebAppStacks(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getProviders().getWebAppStacks(null, com.azure.core.util.Context.NONE);
    }
}
