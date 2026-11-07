
/**
 * Samples for PrivateLinkScopedResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateLinkScopedResourceGet.json
     */
    /**
     * Sample code: Gets private link scoped resource.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getsPrivateLinkScopedResource(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateLinkScopedResources().getWithResponse("MyResourceGroup", "MyPrivateLinkScope",
            "scoped-resource-name", com.azure.core.util.Context.NONE);
    }
}
