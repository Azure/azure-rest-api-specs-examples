import com.azure.core.util.Context;

/** Samples for Provider GetWebAppStacksForLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetWebAppStacksForLocation.json
     */
    /**
     * Sample code: Get Locations Web App Stacks.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getLocationsWebAppStacks(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getProviders()
            .getWebAppStacksForLocation("westus", null, Context.NONE);
    }
}
