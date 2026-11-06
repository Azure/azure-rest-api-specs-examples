
/**
 * Samples for PrivateLinkScopedResources ListByPrivateLinkScope.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateLinkScopedResourceListKindFilter.json
     */
    /**
     * Sample code: Gets list of scoped resources in a private link scope, filtered by kind.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getsListOfScopedResourcesInAPrivateLinkScopeFilteredByKind(
        com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateLinkScopedResources().listByPrivateLinkScope("MyResourceGroup",
            "MyPrivateLinkScope", "Resource", com.azure.core.util.Context.NONE);
    }
}
