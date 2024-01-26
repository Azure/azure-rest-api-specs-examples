
/**
 * Samples for Services ListGloballyEnabledApms.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * Services_ListGloballyEnabledApms.json
     */
    /**
     * Sample code: Services_ListGloballyEnabledApms.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void servicesListGloballyEnabledApms(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getServices()
            .listGloballyEnabledApmsWithResponse("myResourceGroup", "myservice", com.azure.core.util.Context.NONE);
    }
}
