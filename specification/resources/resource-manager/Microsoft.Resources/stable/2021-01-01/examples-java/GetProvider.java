import com.azure.core.util.Context;

/** Samples for Providers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/GetProvider.json
     */
    /**
     * Sample code: Get provider.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getProvider(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .serviceClient()
            .getProviders()
            .getWithResponse("Microsoft.TestRP1", null, Context.NONE);
    }
}
