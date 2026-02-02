
import com.azure.resourcemanager.appservice.fluent.models.CsmPublishingCredentialsPoliciesEntityInner;

/**
 * Samples for WebApps UpdateFtpAllowed.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * UpdatePublishingCredentialsPolicy.json
     */
    /**
     * Sample code: Update FTP Allowed.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateFTPAllowed(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().updateFtpAllowedWithResponse("rg", "testSite",
            new CsmPublishingCredentialsPoliciesEntityInner().withAllow(true), com.azure.core.util.Context.NONE);
    }
}
