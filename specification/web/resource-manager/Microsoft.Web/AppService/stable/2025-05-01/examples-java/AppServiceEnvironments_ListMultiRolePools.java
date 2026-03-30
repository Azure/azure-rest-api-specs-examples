
/**
 * Samples for AppServiceEnvironments ListMultiRolePools.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListMultiRolePools.json
     */
    /**
     * Sample code: Get all multi-role pools.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAllMultiRolePools(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listMultiRolePools("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
