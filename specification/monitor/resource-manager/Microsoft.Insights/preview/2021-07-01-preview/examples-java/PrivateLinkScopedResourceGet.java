
/**
 * Samples for PrivateLinkScopedResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/
     * PrivateLinkScopedResourceGet.json
     */
    /**
     * Sample code: Gets private link scoped resource.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsPrivateLinkScopedResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getPrivateLinkScopedResources().getWithResponse(
            "MyResourceGroup", "MyPrivateLinkScope", "scoped-resource-name", com.azure.core.util.Context.NONE);
    }
}
