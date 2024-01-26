
/**
 * Samples for ApiPortalCustomDomains Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * ApiPortalCustomDomains_Delete.json
     */
    /**
     * Sample code: ApiPortalCustomDomains_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void apiPortalCustomDomainsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getApiPortalCustomDomains().delete("myResourceGroup",
            "myservice", "default", "myDomainName", com.azure.core.util.Context.NONE);
    }
}
