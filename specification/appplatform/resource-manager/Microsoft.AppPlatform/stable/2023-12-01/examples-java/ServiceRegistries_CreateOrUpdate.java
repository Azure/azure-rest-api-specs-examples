
/**
 * Samples for ServiceRegistries CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * ServiceRegistries_CreateOrUpdate.json
     */
    /**
     * Sample code: ServiceRegistries_CreateOrUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void serviceRegistriesCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getServiceRegistries().createOrUpdate("myResourceGroup",
            "myservice", "default", com.azure.core.util.Context.NONE);
    }
}
