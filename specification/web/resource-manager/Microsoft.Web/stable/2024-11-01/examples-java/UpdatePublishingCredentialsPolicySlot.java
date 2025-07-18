
import com.azure.resourcemanager.appservice.fluent.models.CsmPublishingCredentialsPoliciesEntityInner;

/**
 * Samples for WebApps UpdateScmAllowedSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/UpdatePublishingCredentialsPolicySlot
     * .json
     */
    /**
     * Sample code: Update SCM Allowed.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateSCMAllowed(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().updateScmAllowedSlotWithResponse("rg", "testSite",
            "stage", new CsmPublishingCredentialsPoliciesEntityInner().withAllow(true),
            com.azure.core.util.Context.NONE);
    }
}
