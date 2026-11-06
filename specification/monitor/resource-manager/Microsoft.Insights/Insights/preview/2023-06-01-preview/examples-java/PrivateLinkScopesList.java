
/**
 * Samples for PrivateLinkScopes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateLinkScopesList.json
     */
    /**
     * Sample code: PrivateLinkScopesList.json.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void privateLinkScopesListJson(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateLinkScopes().list(com.azure.core.util.Context.NONE);
    }
}
