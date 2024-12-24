
import com.azure.resourcemanager.security.models.RecommendationSupportedClouds;
import com.azure.resourcemanager.security.models.SecurityIssue;
import com.azure.resourcemanager.security.models.SeverityEnum;
import java.util.Arrays;

/**
 * Samples for CustomRecommendations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/CustomRecommendations/
     * PutBySubscriptionCustomRecommendation_example.json
     */
    /**
     * Sample code: Create or update custom recommendation over subscription scope.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void createOrUpdateCustomRecommendationOverSubscriptionScope(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.customRecommendations().define("33e7cc6e-a139-4723-a0e5-76993aee0771")
            .withExistingScope("subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b")
            .withQuery(
                "RawEntityMetadata | where Environment == 'GCP' and Identifiers.Type == 'compute.firewalls' | extend IslogConfigEnabled = tobool(Record.logConfig.enable) | extend HealthStatus = iff(IslogConfigEnabled, 'HEALTHY', 'UNHEALTHY')")
            .withCloudProviders(Arrays.asList(RecommendationSupportedClouds.AWS)).withSeverity(SeverityEnum.MEDIUM)
            .withSecurityIssue(SecurityIssue.VULNERABILITY).withDisplayName("Password Policy")
            .withDescription("organization passwords policy").withRemediationDescription("Change password policy to...")
            .create();
    }
}
