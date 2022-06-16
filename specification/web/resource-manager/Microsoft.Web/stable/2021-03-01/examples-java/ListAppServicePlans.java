import com.azure.core.util.Context;

/** Samples for AppServicePlans List. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListAppServicePlans.json
     */
    /**
     * Sample code: List App Service plans.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAppServicePlans(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServicePlans().list(null, Context.NONE);
    }
}
