
import com.azure.core.util.Context;

/** Samples for PrivateLinkResources ListByPrivateLinkScope. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/
     * PrivateLinkScopePrivateLinkResourceListGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getPrivateLinkResources()
            .listByPrivateLinkScopeWithResponse("MyResourceGroup", "MyPrivateLinkScope", Context.NONE);
    }
}
