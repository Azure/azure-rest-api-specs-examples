
/**
 * Samples for AppServiceEnvironments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_Delete.json
     */
    /**
     * Sample code: Delete an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void deleteAnAppServiceEnvironment(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().delete("test-rg", "test-ase", null,
            com.azure.core.util.Context.NONE);
    }
}
