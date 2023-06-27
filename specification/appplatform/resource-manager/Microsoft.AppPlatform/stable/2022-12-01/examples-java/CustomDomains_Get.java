import com.azure.core.util.Context;

/** Samples for CustomDomains Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/CustomDomains_Get.json
     */
    /**
     * Sample code: CustomDomains_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void customDomainsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getCustomDomains()
            .getWithResponse("myResourceGroup", "myservice", "myapp", "mydomain.com", Context.NONE);
    }
}
