import com.azure.core.util.Context;

/** Samples for Apps Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Apps_Get.json
     */
    /**
     * Sample code: Apps_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void appsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getApps()
            .getWithResponse("myResourceGroup", "myservice", "myapp", null, Context.NONE);
    }
}
