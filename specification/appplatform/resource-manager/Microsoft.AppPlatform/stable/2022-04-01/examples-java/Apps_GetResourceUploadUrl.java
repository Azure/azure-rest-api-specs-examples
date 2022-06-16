import com.azure.core.util.Context;

/** Samples for Apps GetResourceUploadUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Apps_GetResourceUploadUrl.json
     */
    /**
     * Sample code: Apps_GetResourceUploadUrl.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void appsGetResourceUploadUrl(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getApps()
            .getResourceUploadUrlWithResponse("myResourceGroup", "myservice", "myapp", Context.NONE);
    }
}
