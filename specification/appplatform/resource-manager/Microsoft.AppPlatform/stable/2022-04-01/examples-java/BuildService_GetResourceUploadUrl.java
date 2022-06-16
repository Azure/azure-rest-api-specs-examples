import com.azure.core.util.Context;

/** Samples for BuildService GetResourceUploadUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildService_GetResourceUploadUrl.json
     */
    /**
     * Sample code: BuildService_GetResourceUploadUrl.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceGetResourceUploadUrl(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBuildServices()
            .getResourceUploadUrlWithResponse("myResourceGroup", "myservice", "default", Context.NONE);
    }
}
