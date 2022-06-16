import com.azure.core.util.Context;

/** Samples for AppServicePlans Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/DeleteAppServicePlan.json
     */
    /**
     * Sample code: Delete App Service plan.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAppServicePlan(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServicePlans()
            .deleteWithResponse("testrg123", "testsf6141", Context.NONE);
    }
}
