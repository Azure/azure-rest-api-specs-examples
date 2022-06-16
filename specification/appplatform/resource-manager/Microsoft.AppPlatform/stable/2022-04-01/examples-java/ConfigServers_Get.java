import com.azure.core.util.Context;

/** Samples for ConfigServers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/ConfigServers_Get.json
     */
    /**
     * Sample code: ConfigServers_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void configServersGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getConfigServers()
            .getWithResponse("myResourceGroup", "myservice", Context.NONE);
    }
}
