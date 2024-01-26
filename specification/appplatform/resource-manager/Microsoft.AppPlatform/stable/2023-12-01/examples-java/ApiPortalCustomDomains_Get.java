
/**
 * Samples for ApiPortalCustomDomains Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * ApiPortalCustomDomains_Get.json
     */
    /**
     * Sample code: ApiPortalCustomDomains_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void apiPortalCustomDomainsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getApiPortalCustomDomains().getWithResponse("myResourceGroup",
            "myservice", "default", "myDomainName", com.azure.core.util.Context.NONE);
    }
}
