
/**
 * Samples for ProviderResourceTypes List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Resources/stable/2025-04-01/examples/GetProviderResourceTypes.
     * json
     */
    /**
     * Sample code: Get provider resource types.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getProviderResourceTypes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().serviceClient().getProviderResourceTypes()
            .listWithResponse("Microsoft.TestRP", null, com.azure.core.util.Context.NONE);
    }
}
