import com.azure.core.util.Context;

/** Samples for AppServiceEnvironments ListMultiRoleUsages. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/AppServiceEnvironments_ListMultiRoleUsages.json
     */
    /**
     * Sample code: Get usage metrics for a multi-role pool of an App Service Environment.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getUsageMetricsForAMultiRolePoolOfAnAppServiceEnvironment(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServiceEnvironments()
            .listMultiRoleUsages("test-rg", "test-ase", Context.NONE);
    }
}
