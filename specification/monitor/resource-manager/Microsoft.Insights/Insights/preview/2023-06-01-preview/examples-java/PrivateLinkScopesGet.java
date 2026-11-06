
/**
 * Samples for PrivateLinkScopes GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateLinkScopesGet.json
     */
    /**
     * Sample code: PrivateLinkScopeGet.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void privateLinkScopeGet(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateLinkScopes().getByResourceGroupWithResponse("my-resource-group",
            "my-privatelinkscope", com.azure.core.util.Context.NONE);
    }
}
