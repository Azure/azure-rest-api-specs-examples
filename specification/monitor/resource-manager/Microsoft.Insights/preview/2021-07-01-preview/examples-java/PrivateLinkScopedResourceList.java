
import com.azure.core.util.Context;

/** Samples for PrivateLinkScopedResources ListByPrivateLinkScope. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/
     * PrivateLinkScopedResourceList.json
     */
    /**
     * Sample code: Gets list of scoped resources in a private link scope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getsListOfScopedResourcesInAPrivateLinkScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getPrivateLinkScopedResources()
            .listByPrivateLinkScope("MyResourceGroup", "MyPrivateLinkScope", Context.NONE);
    }
}
