/** Samples for ApiManagementService MigrateToStv2. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementServiceMigrateToStv2.json
     */
    /**
     * Sample code: ApiManagementMigrateService.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementMigrateService(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiManagementServices().migrateToStv2("rg1", "apimService1", com.azure.core.util.Context.NONE);
    }
}
