
/**
 * Samples for Provider GetFunctionAppStacks.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetFunctionAppStacks.json
     */
    /**
     * Sample code: Get Function App Stacks.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFunctionAppStacks(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getProviders().getFunctionAppStacks(null,
            com.azure.core.util.Context.NONE);
    }
}
