
/**
 * Samples for CustomDomains Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/CustomDomains_Delete.
     * json
     */
    /**
     * Sample code: CustomDomains_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void customDomainsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getCustomDomains().delete("myResourceGroup", "myservice",
            "myapp", "mydomain.com", com.azure.core.util.Context.NONE);
    }
}
