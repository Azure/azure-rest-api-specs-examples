
/**
 * Samples for WebApps ListBasicPublishingCredentialsPolicies.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * ListPublishingCredentialsPolicies.json
     */
    /**
     * Sample code: List Publishing Credentials Policies.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPublishingCredentialsPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().listBasicPublishingCredentialsPolicies("testrg123",
            "testsite", com.azure.core.util.Context.NONE);
    }
}
