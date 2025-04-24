
import com.azure.resourcemanager.apimanagement.fluent.models.PortalSignupSettingsInner;
import com.azure.resourcemanager.apimanagement.models.TermsOfServiceProperties;

/**
 * Samples for SignUpSettings Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementPortalSettingsUpdateSignUp.json
     */
    /**
     * Sample code: ApiManagementPortalSettingsUpdateSignUp.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementPortalSettingsUpdateSignUp(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.signUpSettings().updateWithResponse("rg1", "apimService1", "*",
            new PortalSignupSettingsInner().withEnabled(true).withTermsOfService(new TermsOfServiceProperties()
                .withText("Terms of service text.").withEnabled(true).withConsentRequired(true)),
            com.azure.core.util.Context.NONE);
    }
}
