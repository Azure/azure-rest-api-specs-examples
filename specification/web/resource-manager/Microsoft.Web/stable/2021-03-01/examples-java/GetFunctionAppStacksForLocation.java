import com.azure.core.util.Context;

/** Samples for Provider GetFunctionAppStacksForLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetFunctionAppStacksForLocation.json
     */
    /**
     * Sample code: Get Locations Function App Stacks.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getLocationsFunctionAppStacks(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getProviders()
            .getFunctionAppStacksForLocation("westus", null, Context.NONE);
    }
}
