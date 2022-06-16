import com.azure.core.util.Context;

/** Samples for ConfigurationServices Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/ConfigurationServices_Get.json
     */
    /**
     * Sample code: ConfigurationServices_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void configurationServicesGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getConfigurationServices()
            .getWithResponse("myResourceGroup", "myservice", "default", Context.NONE);
    }
}
