
import com.azure.resourcemanager.appservice.fluent.models.CsmPublishingCredentialsPoliciesEntityInner;

/**
 * Samples for WebApps UpdateFtpAllowedSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/UpdatePublishingCredentialsPolicySlot
     * .json
     */
    /**
     * Sample code: Update FTP Allowed.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateFTPAllowed(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().updateFtpAllowedSlotWithResponse("rg", "testSite",
            "stage", new CsmPublishingCredentialsPoliciesEntityInner().withAllow(true),
            com.azure.core.util.Context.NONE);
    }
}
