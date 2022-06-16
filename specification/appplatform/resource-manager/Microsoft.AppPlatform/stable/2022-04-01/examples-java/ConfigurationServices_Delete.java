import com.azure.core.util.Context;

/** Samples for ConfigurationServices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/ConfigurationServices_Delete.json
     */
    /**
     * Sample code: ConfigurationServices_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void configurationServicesDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getConfigurationServices()
            .delete("myResourceGroup", "myservice", "default", Context.NONE);
    }
}
