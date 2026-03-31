
/**
 * Samples for AppServiceEnvironments Resume.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_Resume.json
     */
    /**
     * Sample code: Resume an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void resumeAnAppServiceEnvironment(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().resume("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
