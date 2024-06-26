
/**
 * Samples for Services List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * Services_ListBySubscription.json
     */
    /**
     * Sample code: Services_ListBySubscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void servicesListBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getServices().list(com.azure.core.util.Context.NONE);
    }
}
