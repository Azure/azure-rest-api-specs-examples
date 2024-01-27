
/**
 * Samples for DevToolPortals Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/DevToolPortal_Delete.
     * json
     */
    /**
     * Sample code: DevToolPortals_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void devToolPortalsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getDevToolPortals().delete("myResourceGroup", "myservice",
            "default", com.azure.core.util.Context.NONE);
    }
}
