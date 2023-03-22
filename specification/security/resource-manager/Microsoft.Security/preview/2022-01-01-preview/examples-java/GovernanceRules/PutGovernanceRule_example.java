import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.security.models.GovernanceRuleEmailNotification;
import com.azure.resourcemanager.security.models.GovernanceRuleOwnerSource;
import com.azure.resourcemanager.security.models.GovernanceRuleOwnerSourceType;
import com.azure.resourcemanager.security.models.GovernanceRuleSourceResourceType;
import com.azure.resourcemanager.security.models.GovernanceRuleType;
import java.io.IOException;
import java.util.Arrays;

/** Samples for GovernanceRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/PutGovernanceRule_example.json
     */
    /**
     * Sample code: Create or update governance rule over subscription scope.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void createOrUpdateGovernanceRuleOverSubscriptionScope(
        com.azure.resourcemanager.security.SecurityManager manager) throws IOException {
        manager
            .governanceRules()
            .define("ad9a8e26-29d9-4829-bb30-e597a58cdbb8")
            .withExistingScope("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23")
            .withDisplayName("Admin's rule")
            .withDescription("A rule for critical recommendations")
            .withRemediationTimeframe("7.00:00:00")
            .withIsGracePeriod(true)
            .withRulePriority(200)
            .withIsDisabled(false)
            .withRuleType(GovernanceRuleType.INTEGRATED)
            .withSourceResourceType(GovernanceRuleSourceResourceType.ASSESSMENTS)
            .withConditionSets(
                Arrays
                    .asList(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize(
                                "{\"conditions\":[{\"operator\":\"In\",\"property\":\"$.AssessmentKey\",\"value\":\"[\\\"b1cd27e0-4ecc-4246-939f-49c426d9d72f\\\","
                                    + " \\\"fe83f80b-073d-4ccf-93d9-6797eb870201\\\"]\"}]}",
                                Object.class,
                                SerializerEncoding.JSON)))
            .withOwnerSource(
                new GovernanceRuleOwnerSource()
                    .withType(GovernanceRuleOwnerSourceType.MANUALLY)
                    .withValue("user@contoso.com"))
            .withGovernanceEmailNotification(
                new GovernanceRuleEmailNotification()
                    .withDisableManagerEmailNotification(false)
                    .withDisableOwnerEmailNotification(false))
            .create();
    }
}
