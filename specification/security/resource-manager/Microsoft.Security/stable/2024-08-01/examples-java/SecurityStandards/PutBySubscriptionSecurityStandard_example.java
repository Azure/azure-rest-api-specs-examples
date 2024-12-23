
import com.azure.resourcemanager.security.models.PartialAssessmentProperties;
import com.azure.resourcemanager.security.models.StandardSupportedCloud;
import java.util.Arrays;

/**
 * Samples for SecurityStandards CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/SecurityStandards/
     * PutBySubscriptionSecurityStandard_example.json
     */
    /**
     * Sample code: Create or update security standard over subscription scope.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void createOrUpdateSecurityStandardOverSubscriptionScope(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityStandards().define("8bb8be0a-6010-4789-812f-e4d661c4ed0e")
            .withExistingScope("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23")
            .withDisplayName("Azure Test Security Standard 1")
            .withDescription("description of Azure Test Security Standard 1")
            .withAssessments(Arrays.asList(new PartialAssessmentProperties().withAssessmentKey("fakeTokenPlaceholder"),
                new PartialAssessmentProperties().withAssessmentKey("fakeTokenPlaceholder")))
            .withCloudProviders(Arrays.asList(StandardSupportedCloud.GCP))
            .withPolicySetDefinitionId(
                "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Authorization/policySetDefinitions/patchorchestration-applicationversions")
            .create();
    }
}
