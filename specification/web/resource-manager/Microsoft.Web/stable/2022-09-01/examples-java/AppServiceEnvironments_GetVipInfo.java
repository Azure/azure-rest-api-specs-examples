/** Samples for AppServiceEnvironments GetVipInfo. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/AppServiceEnvironments_GetVipInfo.json
     */
    /**
     * Sample code: Get IP addresses assigned to an App Service Environment.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getIPAddressesAssignedToAnAppServiceEnvironment(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServiceEnvironments()
            .getVipInfoWithResponse("test-rg", "test-ase", com.azure.core.util.Context.NONE);
    }
}
