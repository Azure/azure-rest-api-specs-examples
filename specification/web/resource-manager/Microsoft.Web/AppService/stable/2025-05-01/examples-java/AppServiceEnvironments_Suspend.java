
/**
 * Samples for AppServiceEnvironments Suspend.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_Suspend.json
     */
    /**
     * Sample code: Suspend an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void suspendAnAppServiceEnvironment(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().suspend("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
