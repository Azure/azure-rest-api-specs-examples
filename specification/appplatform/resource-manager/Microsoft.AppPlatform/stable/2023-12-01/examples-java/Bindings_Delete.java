
/**
 * Samples for Bindings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Bindings_Delete.json
     */
    /**
     * Sample code: Bindings_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void bindingsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getBindings().delete("myResourceGroup", "myservice", "myapp",
            "mybinding", com.azure.core.util.Context.NONE);
    }
}
