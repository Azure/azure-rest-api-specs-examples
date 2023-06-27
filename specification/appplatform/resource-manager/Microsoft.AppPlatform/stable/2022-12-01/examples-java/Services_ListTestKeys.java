import com.azure.core.util.Context;

/** Samples for Services ListTestKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Services_ListTestKeys.json
     */
    /**
     * Sample code: Services_ListTestKeys.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void servicesListTestKeys(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getServices()
            .listTestKeysWithResponse("myResourceGroup", "myservice", Context.NONE);
    }
}
