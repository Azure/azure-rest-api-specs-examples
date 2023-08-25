import com.azure.resourcemanager.apimanagement.models.ConfigurationIdName;
import com.azure.resourcemanager.apimanagement.models.DeployConfigurationParameters;

/** Samples for TenantConfiguration Validate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementTenantConfigurationValidate.json
     */
    /**
     * Sample code: ApiManagementTenantConfigurationValidate.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementTenantConfigurationValidate(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tenantConfigurations()
            .validate(
                "rg1",
                "apimService1",
                ConfigurationIdName.CONFIGURATION,
                new DeployConfigurationParameters().withBranch("master"),
                com.azure.core.util.Context.NONE);
    }
}
