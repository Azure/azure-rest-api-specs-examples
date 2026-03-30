
/**
 * Samples for AppServiceEnvironments GetMultiRolePool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_GetMultiRolePool.json
     */
    /**
     * Sample code: Get properties of a multi-role pool.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getPropertiesOfAMultiRolePool(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().getMultiRolePoolWithResponse("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
