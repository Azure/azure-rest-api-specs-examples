
import com.azure.resourcemanager.appservice.fluent.models.CsmPublishingCredentialsPoliciesEntityInner;

/**
 * Samples for WebApps UpdateScmAllowed.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/UpdatePublishingCredentialsPolicy.json
     */
    /**
     * Sample code: Update SCM Allowed.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void updateSCMAllowed(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().updateScmAllowedWithResponse("rg", "testSite",
            new CsmPublishingCredentialsPoliciesEntityInner().withAllow(true), com.azure.core.util.Context.NONE);
    }
}
