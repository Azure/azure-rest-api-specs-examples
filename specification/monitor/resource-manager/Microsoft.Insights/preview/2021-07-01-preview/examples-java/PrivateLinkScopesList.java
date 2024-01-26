
import com.azure.core.util.Context;

/** Samples for PrivateLinkScopes List. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/
     * PrivateLinkScopesList.json
     */
    /**
     * Sample code: PrivateLinkScopesList.json.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void privateLinkScopesListJson(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getPrivateLinkScopes().list(Context.NONE);
    }
}
