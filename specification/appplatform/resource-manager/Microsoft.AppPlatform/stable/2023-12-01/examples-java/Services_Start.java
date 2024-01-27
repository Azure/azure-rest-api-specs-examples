
/**
 * Samples for Services Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Services_Start.json
     */
    /**
     * Sample code: Services_Start.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void servicesStart(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getServices().start("myResourceGroup", "myservice",
            com.azure.core.util.Context.NONE);
    }
}
