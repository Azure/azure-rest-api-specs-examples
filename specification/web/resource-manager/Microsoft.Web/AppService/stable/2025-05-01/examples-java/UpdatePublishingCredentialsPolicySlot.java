
import com.azure.resourcemanager.appservice.fluent.models.CsmPublishingCredentialsPoliciesEntityInner;

/**
 * Samples for WebApps UpdateScmAllowedSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/UpdatePublishingCredentialsPolicySlot.json
     */
    /**
     * Sample code: Update SCM Allowed.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void updateSCMAllowed(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().updateScmAllowedSlotWithResponse("rg", "testSite", "stage",
            new CsmPublishingCredentialsPoliciesEntityInner().withAllow(true), com.azure.core.util.Context.NONE);
    }
}
