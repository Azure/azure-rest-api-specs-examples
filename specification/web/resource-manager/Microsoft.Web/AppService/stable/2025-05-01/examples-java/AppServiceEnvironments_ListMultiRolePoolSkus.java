
/**
 * Samples for AppServiceEnvironments ListMultiRolePoolSkus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListMultiRolePoolSkus.json
     */
    /**
     * Sample code: Get available SKUs for scaling a multi-role pool.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getAvailableSKUsForScalingAMultiRolePool(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listMultiRolePoolSkus("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
