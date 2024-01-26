
/**
 * Samples for Services ListSupportedApmTypes.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * Services_ListSupportedApmTypes.json
     */
    /**
     * Sample code: Services_ListSupportedApmTypes.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void servicesListSupportedApmTypes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getServices().listSupportedApmTypes("myResourceGroup",
            "myservice", com.azure.core.util.Context.NONE);
    }
}
