
/**
 * Samples for GatewayCustomDomains Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * GatewayCustomDomains_Get.json
     */
    /**
     * Sample code: GatewayCustomDomains_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gatewayCustomDomainsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getGatewayCustomDomains().getWithResponse("myResourceGroup",
            "myservice", "default", "myDomainName", com.azure.core.util.Context.NONE);
    }
}
