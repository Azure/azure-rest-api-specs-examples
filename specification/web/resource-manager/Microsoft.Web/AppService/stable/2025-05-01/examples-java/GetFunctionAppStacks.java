
/**
 * Samples for Provider GetFunctionAppStacks.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetFunctionAppStacks.json
     */
    /**
     * Sample code: Get Function App Stacks.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getFunctionAppStacks(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getProviders().getFunctionAppStacks(null, com.azure.core.util.Context.NONE);
    }
}
