import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.ConfigurationIdName;
import com.azure.resourcemanager.apimanagement.models.DeployConfigurationParameters;

/** Samples for TenantConfiguration Deploy. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementTenantConfigurationDeploy.json
     */
    /**
     * Sample code: ApiManagementTenantConfigurationDeploy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementTenantConfigurationDeploy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tenantConfigurations()
            .deploy(
                "rg1",
                "apimService1",
                ConfigurationIdName.CONFIGURATION,
                new DeployConfigurationParameters().withBranch("master"),
                Context.NONE);
    }
}
