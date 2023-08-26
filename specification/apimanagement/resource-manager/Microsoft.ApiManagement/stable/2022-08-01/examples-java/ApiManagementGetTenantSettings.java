import com.azure.resourcemanager.apimanagement.models.SettingsTypeName;

/** Samples for TenantSettings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetTenantSettings.json
     */
    /**
     * Sample code: ApiManagementGetTenantSettings.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetTenantSettings(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tenantSettings()
            .getWithResponse("rg1", "apimService1", SettingsTypeName.PUBLIC, com.azure.core.util.Context.NONE);
    }
}
