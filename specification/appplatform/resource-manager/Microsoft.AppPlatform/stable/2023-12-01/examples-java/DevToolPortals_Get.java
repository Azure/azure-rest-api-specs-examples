
/**
 * Samples for DevToolPortals Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/DevToolPortals_Get.
     * json
     */
    /**
     * Sample code: DevToolPortals_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void devToolPortalsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getDevToolPortals().getWithResponse("myResourceGroup",
            "myservice", "default", com.azure.core.util.Context.NONE);
    }
}
