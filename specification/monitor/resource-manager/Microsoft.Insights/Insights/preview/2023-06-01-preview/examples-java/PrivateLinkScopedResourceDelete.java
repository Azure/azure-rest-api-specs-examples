
/**
 * Samples for PrivateLinkScopedResources Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/PrivateLinkScopedResourceDelete.json
     */
    /**
     * Sample code: Deletes a scoped resource with a given name.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void deletesAScopedResourceWithAGivenName(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateLinkScopedResources().delete("MyResourceGroup", "MyPrivateLinkScope",
            "scoped-resource-name", com.azure.core.util.Context.NONE);
    }
}
