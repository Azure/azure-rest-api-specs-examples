import com.azure.core.util.Context;

/** Samples for ServiceRegistries List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/ServiceRegistries_List.json
     */
    /**
     * Sample code: ServiceRegistries_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void serviceRegistriesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getServiceRegistries()
            .list("myResourceGroup", "myservice", Context.NONE);
    }
}
