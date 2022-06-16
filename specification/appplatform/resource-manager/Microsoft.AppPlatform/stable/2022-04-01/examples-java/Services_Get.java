import com.azure.core.util.Context;

/** Samples for Services GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Services_Get.json
     */
    /**
     * Sample code: Services_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void servicesGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getServices()
            .getByResourceGroupWithResponse("myResourceGroup", "myservice", Context.NONE);
    }
}
