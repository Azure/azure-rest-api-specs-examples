import com.azure.core.util.Context;

/** Samples for Providers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/GetProviders.json
     */
    /**
     * Sample code: Get providers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getProviders(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().serviceClient().getProviders().list(null, null, Context.NONE);
    }
}
