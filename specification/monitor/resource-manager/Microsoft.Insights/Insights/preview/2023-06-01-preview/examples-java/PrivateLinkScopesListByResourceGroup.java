
/**
 * Samples for PrivateLinkScopes ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateLinkScopesListByResourceGroup.json
     */
    /**
     * Sample code: PrivateLinkScopeListByResourceGroup.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void privateLinkScopeListByResourceGroup(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateLinkScopes().listByResourceGroup("my-resource-group",
            com.azure.core.util.Context.NONE);
    }
}
