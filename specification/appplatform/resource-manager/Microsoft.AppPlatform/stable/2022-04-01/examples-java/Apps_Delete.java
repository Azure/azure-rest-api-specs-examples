import com.azure.core.util.Context;

/** Samples for Apps Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Apps_Delete.json
     */
    /**
     * Sample code: Apps_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void appsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getApps()
            .delete("myResourceGroup", "myservice", "myapp", Context.NONE);
    }
}
