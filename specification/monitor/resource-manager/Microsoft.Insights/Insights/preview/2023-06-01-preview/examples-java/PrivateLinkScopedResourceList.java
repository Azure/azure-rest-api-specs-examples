
/**
 * Samples for PrivateLinkScopedResources ListByPrivateLinkScope.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateLinkScopedResourceList.json
     */
    /**
     * Sample code: Gets list of scoped resources in a private link scope.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        getsListOfScopedResourcesInAPrivateLinkScope(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateLinkScopedResources().listByPrivateLinkScope("MyResourceGroup",
            "MyPrivateLinkScope", null, com.azure.core.util.Context.NONE);
    }
}
