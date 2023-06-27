import com.azure.core.util.Context;

/** Samples for ServiceRegistries Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/ServiceRegistries_Get.json
     */
    /**
     * Sample code: ServiceRegistries_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void serviceRegistriesGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getServiceRegistries()
            .getWithResponse("myResourceGroup", "myservice", "default", Context.NONE);
    }
}
