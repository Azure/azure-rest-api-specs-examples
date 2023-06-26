import com.azure.core.util.Context;

/** Samples for ConfigurationServices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/ConfigurationServices_List.json
     */
    /**
     * Sample code: ConfigurationServices_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void configurationServicesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getConfigurationServices()
            .list("myResourceGroup", "myservice", Context.NONE);
    }
}
