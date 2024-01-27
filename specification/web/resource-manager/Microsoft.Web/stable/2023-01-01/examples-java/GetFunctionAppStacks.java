
/**
 * Samples for Provider GetFunctionAppStacks.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/GetFunctionAppStacks.json
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
