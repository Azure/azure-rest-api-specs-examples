
/**
 * Samples for PrivateLinkScopes GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/
     * PrivateLinkScopesGet.json
     */
    /**
     * Sample code: PrivateLinkScopeGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void privateLinkScopeGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getPrivateLinkScopes().getByResourceGroupWithResponse(
            "my-resource-group", "my-privatelinkscope", com.azure.core.util.Context.NONE);
    }
}
