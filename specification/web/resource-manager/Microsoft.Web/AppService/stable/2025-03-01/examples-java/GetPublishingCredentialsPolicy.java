
/**
 * Samples for WebApps GetFtpAllowed.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetPublishingCredentialsPolicy.json
     */
    /**
     * Sample code: Get FTP Allowed.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFTPAllowed(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getFtpAllowedWithResponse("rg", "testSite",
            com.azure.core.util.Context.NONE);
    }
}
