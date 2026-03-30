
/**
 * Samples for AppServiceEnvironments GetVipInfo.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_GetVipInfo.json
     */
    /**
     * Sample code: Get IP addresses assigned to an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getIPAddressesAssignedToAnAppServiceEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().getVipInfoWithResponse("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
