
import com.azure.resourcemanager.apimanagement.models.ConfigurationIdName;
import com.azure.resourcemanager.apimanagement.models.SaveConfigurationParameter;

/**
 * Samples for TenantConfiguration Save.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementTenantConfigurationSave.json
     */
    /**
     * Sample code: ApiManagementTenantConfigurationSave.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementTenantConfigurationSave(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tenantConfigurations().save("rg1", "apimService1", ConfigurationIdName.CONFIGURATION,
            new SaveConfigurationParameter().withBranch("master"), com.azure.core.util.Context.NONE);
    }
}
