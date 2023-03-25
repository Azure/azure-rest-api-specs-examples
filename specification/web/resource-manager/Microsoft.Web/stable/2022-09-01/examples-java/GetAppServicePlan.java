/** Samples for AppServicePlans GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/GetAppServicePlan.json
     */
    /**
     * Sample code: Get App Service plan.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppServicePlan(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServicePlans()
            .getByResourceGroupWithResponse("testrg123", "testsf6141", com.azure.core.util.Context.NONE);
    }
}
