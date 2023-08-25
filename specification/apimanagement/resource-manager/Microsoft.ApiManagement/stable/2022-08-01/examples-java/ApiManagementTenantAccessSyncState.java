import com.azure.resourcemanager.apimanagement.models.ConfigurationIdName;

/** Samples for TenantConfiguration GetSyncState. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementTenantAccessSyncState.json
     */
    /**
     * Sample code: ApiManagementTenantAccessSyncState.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementTenantAccessSyncState(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tenantConfigurations()
            .getSyncStateWithResponse(
                "rg1", "apimService1", ConfigurationIdName.CONFIGURATION, com.azure.core.util.Context.NONE);
    }
}
