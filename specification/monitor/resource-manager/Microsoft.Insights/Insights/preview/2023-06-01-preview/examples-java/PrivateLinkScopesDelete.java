
/**
 * Samples for PrivateLinkScopes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateLinkScopesDelete.json
     */
    /**
     * Sample code: PrivateLinkScopesDelete.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void privateLinkScopesDelete(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateLinkScopes().delete("my-resource-group", "my-privatelinkscope",
            com.azure.core.util.Context.NONE);
    }
}
