Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.AlertDetail;
import com.azure.resourcemanager.securityinsights.models.AlertDetailsOverride;
import com.azure.resourcemanager.securityinsights.models.AlertSeverity;
import com.azure.resourcemanager.securityinsights.models.AttackTactic;
import com.azure.resourcemanager.securityinsights.models.EntityMapping;
import com.azure.resourcemanager.securityinsights.models.EntityMappingType;
import com.azure.resourcemanager.securityinsights.models.EventGroupingAggregationKind;
import com.azure.resourcemanager.securityinsights.models.EventGroupingSettings;
import com.azure.resourcemanager.securityinsights.models.FieldMapping;
import com.azure.resourcemanager.securityinsights.models.GroupingConfiguration;
import com.azure.resourcemanager.securityinsights.models.IncidentConfiguration;
import com.azure.resourcemanager.securityinsights.models.MatchingMethod;
import com.azure.resourcemanager.securityinsights.models.ScheduledAlertRule;
import com.azure.resourcemanager.securityinsights.models.TriggerOperator;
import java.time.Duration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for AlertRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/alertRules/CreateScheduledAlertRule.json
     */
    /**
     * Sample code: Creates or updates a Scheduled alert rule.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAScheduledAlertRule(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .alertRules()
            .createOrUpdateWithResponse(
                "myRg",
                "myWorkspace",
                "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
                new ScheduledAlertRule()
                    .withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"")
                    .withDescription("An example for a scheduled rule")
                    .withDisplayName("My scheduled rule")
                    .withEnabled(true)
                    .withSuppressionDuration(Duration.parse("PT1H"))
                    .withSuppressionEnabled(false)
                    .withTactics(Arrays.asList(AttackTactic.PERSISTENCE, AttackTactic.LATERAL_MOVEMENT))
                    .withTechniques(Arrays.asList("T1037", "T1021"))
                    .withIncidentConfiguration(
                        new IncidentConfiguration()
                            .withCreateIncident(true)
                            .withGroupingConfiguration(
                                new GroupingConfiguration()
                                    .withEnabled(true)
                                    .withReopenClosedIncident(false)
                                    .withLookbackDuration(Duration.parse("PT5H"))
                                    .withMatchingMethod(MatchingMethod.SELECTED)
                                    .withGroupByEntities(Arrays.asList(EntityMappingType.HOST))
                                    .withGroupByAlertDetails(Arrays.asList(AlertDetail.DISPLAY_NAME))
                                    .withGroupByCustomDetails(
                                        Arrays.asList("OperatingSystemType", "OperatingSystemName"))))
                    .withQuery("Heartbeat")
                    .withQueryFrequency(Duration.parse("PT1H"))
                    .withQueryPeriod(Duration.parse("P2DT1H30M"))
                    .withSeverity(AlertSeverity.HIGH)
                    .withTriggerOperator(TriggerOperator.GREATER_THAN)
                    .withTriggerThreshold(0)
                    .withEventGroupingSettings(
                        new EventGroupingSettings().withAggregationKind(EventGroupingAggregationKind.ALERT_PER_RESULT))
                    .withCustomDetails(mapOf("OperatingSystemName", "OSName", "OperatingSystemType", "OSType"))
                    .withEntityMappings(
                        Arrays
                            .asList(
                                new EntityMapping()
                                    .withEntityType(EntityMappingType.HOST)
                                    .withFieldMappings(
                                        Arrays
                                            .asList(
                                                new FieldMapping()
                                                    .withIdentifier("FullName")
                                                    .withColumnName("Computer"))),
                                new EntityMapping()
                                    .withEntityType(EntityMappingType.IP)
                                    .withFieldMappings(
                                        Arrays
                                            .asList(
                                                new FieldMapping()
                                                    .withIdentifier("Address")
                                                    .withColumnName("ComputerIP")))))
                    .withAlertDetailsOverride(
                        new AlertDetailsOverride()
                            .withAlertDisplayNameFormat("Alert from {{Computer}}")
                            .withAlertDescriptionFormat("Suspicious activity was made by {{ComputerIP}}")),
                Context.NONE);
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
```
