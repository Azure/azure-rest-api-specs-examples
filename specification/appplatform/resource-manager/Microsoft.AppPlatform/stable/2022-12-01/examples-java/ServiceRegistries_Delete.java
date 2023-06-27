import com.azure.core.util.Context;

/** Samples for ServiceRegistries Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/ServiceRegistries_Delete.json
     */
    /**
     * Sample code: ServiceRegistries_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void serviceRegistriesDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getServiceRegistries()
            .delete("myResourceGroup", "myservice", "default", Context.NONE);
    }
}
