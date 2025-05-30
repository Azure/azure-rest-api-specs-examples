
import com.azure.resourcemanager.apimanagement.models.PortalConfigContract;
import com.azure.resourcemanager.apimanagement.models.PortalConfigCorsProperties;
import com.azure.resourcemanager.apimanagement.models.PortalConfigCspProperties;
import com.azure.resourcemanager.apimanagement.models.PortalConfigDelegationProperties;
import com.azure.resourcemanager.apimanagement.models.PortalConfigPropertiesSignin;
import com.azure.resourcemanager.apimanagement.models.PortalConfigPropertiesSignup;
import com.azure.resourcemanager.apimanagement.models.PortalConfigTermsOfServiceProperties;
import com.azure.resourcemanager.apimanagement.models.PortalSettingsCspMode;
import java.util.Arrays;

/**
 * Samples for PortalConfig Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdatePortalConfig.json
     */
    /**
     * Sample code: ApiManagementUpdatePortalConfig.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdatePortalConfig(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        PortalConfigContract resource = manager.portalConfigs()
            .getWithResponse("rg1", "apimService1", "default", com.azure.core.util.Context.NONE).getValue();
        resource.update().withEnableBasicAuth(true).withSignin(new PortalConfigPropertiesSignin().withRequire(false))
            .withSignup(new PortalConfigPropertiesSignup().withTermsOfService(new PortalConfigTermsOfServiceProperties()
                .withText("I agree to the service terms and conditions.").withRequireConsent(false)))
            .withDelegation(
                new PortalConfigDelegationProperties().withDelegateRegistration(false).withDelegateSubscription(false))
            .withCors(new PortalConfigCorsProperties().withAllowedOrigins(Arrays.asList("https://contoso.com")))
            .withCsp(new PortalConfigCspProperties().withMode(PortalSettingsCspMode.REPORT_ONLY)
                .withReportUri(Arrays.asList("https://report.contoso.com"))
                .withAllowedSources(Arrays.asList("*.contoso.com")))
            .withIfMatch("*").apply();
    }
}
