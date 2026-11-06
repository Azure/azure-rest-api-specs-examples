
import com.azure.resourcemanager.monitor.fluent.models.ScopedResourceInner;
import com.azure.resourcemanager.monitor.models.ScopedResourceKind;

/**
 * Samples for PrivateLinkScopedResources CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateLinkScopedResourceUpdatePlatformMetrics.json
     */
    /**
     * Sample code: Update a scoped platform metrics subscription in a private link scope.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void updateAScopedPlatformMetricsSubscriptionInAPrivateLinkScope(
        com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateLinkScopedResources().createOrUpdate("MyResourceGroup", "MyPrivateLinkScope",
            "scoped-resource-name",
            new ScopedResourceInner().withKind(ScopedResourceKind.fromString("PlatformMetrics"))
                .withLinkedResourceId("/subscriptions/00000000-1111-2222-3333-444444444444")
                .withSubscriptionLocation("eastus"),
            com.azure.core.util.Context.NONE);
    }
}
