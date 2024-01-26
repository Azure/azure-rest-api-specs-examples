
/**
 * Samples for Services Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Services_Delete.json
     */
    /**
     * Sample code: Services_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void servicesDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getServices().delete("myResourceGroup", "myservice",
            com.azure.core.util.Context.NONE);
    }
}
