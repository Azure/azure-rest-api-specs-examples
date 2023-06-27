import com.azure.core.util.Context;

/** Samples for Bindings List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Bindings_List.json
     */
    /**
     * Sample code: Bindings_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void bindingsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBindings()
            .list("myResourceGroup", "myservice", "myapp", Context.NONE);
    }
}
