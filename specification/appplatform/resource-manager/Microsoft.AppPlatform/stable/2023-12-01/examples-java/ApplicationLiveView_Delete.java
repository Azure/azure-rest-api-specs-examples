
/**
 * Samples for ApplicationLiveViews Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * ApplicationLiveView_Delete.json
     */
    /**
     * Sample code: ApplicationLiveViews_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void applicationLiveViewsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getApplicationLiveViews().delete("myResourceGroup",
            "myservice", "default", com.azure.core.util.Context.NONE);
    }
}
