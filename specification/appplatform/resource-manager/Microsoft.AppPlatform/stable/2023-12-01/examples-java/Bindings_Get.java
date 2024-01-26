
/**
 * Samples for Bindings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Bindings_Get.json
     */
    /**
     * Sample code: Bindings_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void bindingsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getBindings().getWithResponse("myResourceGroup", "myservice",
            "myapp", "mybinding", com.azure.core.util.Context.NONE);
    }
}
