
/**
 * Samples for AppServiceEnvironments Reboot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_Reboot.json
     */
    /**
     * Sample code: Reboot all machines in an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        rebootAllMachinesInAnAppServiceEnvironment(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().rebootWithResponse("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
