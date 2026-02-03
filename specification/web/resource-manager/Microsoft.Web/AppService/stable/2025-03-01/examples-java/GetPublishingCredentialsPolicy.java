
/**
 * Samples for WebApps GetScmAllowed.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetPublishingCredentialsPolicy.json
     */
    /**
     * Sample code: Get SCM Allowed.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSCMAllowed(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getScmAllowedWithResponse("rg", "testSite",
            com.azure.core.util.Context.NONE);
    }
}
