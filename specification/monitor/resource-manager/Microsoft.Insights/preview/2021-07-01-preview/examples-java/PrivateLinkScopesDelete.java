
import com.azure.core.util.Context;

/** Samples for PrivateLinkScopes Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/
     * PrivateLinkScopesDelete.json
     */
    /**
     * Sample code: PrivateLinkScopesDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void privateLinkScopesDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getPrivateLinkScopes().delete("my-resource-group",
            "my-privatelinkscope", Context.NONE);
    }
}
