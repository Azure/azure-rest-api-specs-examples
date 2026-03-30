
import com.azure.resourcemanager.appservice.fluent.models.CsmPublishingCredentialsPoliciesEntityInner;

/**
 * Samples for WebApps UpdateFtpAllowed.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/UpdatePublishingCredentialsPolicy_UpdateFtpAllowed.json
     */
    /**
     * Sample code: Update FTP Allowed.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void updateFTPAllowed(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().updateFtpAllowedWithResponse("rg", "testSite",
            new CsmPublishingCredentialsPoliciesEntityInner().withAllow(true), com.azure.core.util.Context.NONE);
    }
}
