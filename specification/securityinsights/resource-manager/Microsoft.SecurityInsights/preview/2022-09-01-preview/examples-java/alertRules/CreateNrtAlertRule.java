import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.AlertSeverity;
import com.azure.resourcemanager.securityinsights.models.AttackTactic;
import com.azure.resourcemanager.securityinsights.models.EntityMappingType;
import com.azure.resourcemanager.securityinsights.models.EventGroupingAggregationKind;
import com.azure.resourcemanager.securityinsights.models.EventGroupingSettings;
import com.azure.resourcemanager.securityinsights.models.GroupingConfiguration;
import com.azure.resourcemanager.securityinsights.models.IncidentConfiguration;
import com.azure.resourcemanager.securityinsights.models.MatchingMethod;
import com.azure.resourcemanager.securityinsights.models.NrtAlertRule;
import java.time.Duration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for AlertRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/CreateNrtAlertRule.json
     */
    /**
     * Sample code: Creates or updates a Nrt alert rule.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesANrtAlertRule(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .alertRules()
            .createOrUpdateWithResponse(
                "myRg",
                "myWorkspace",
                "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
                new NrtAlertRule()
                    .withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"")
                    .withDescription("")
                    .withQuery(
                        "ProtectionStatus | extend HostCustomEntity = Computer | extend IPCustomEntity ="
                            + " ComputerIP_Hidden")
                    .withTactics(Arrays.asList(AttackTactic.PERSISTENCE, AttackTactic.LATERAL_MOVEMENT))
                    .withTechniques(Arrays.asList("T1037", "T1021"))
                    .withDisplayName("Rule2")
                    .withEnabled(true)
                    .withSuppressionDuration(Duration.parse("PT1H"))
                    .withSuppressionEnabled(false)
                    .withSeverity(AlertSeverity.HIGH)
                    .withIncidentConfiguration(
                        new IncidentConfiguration()
                            .withCreateIncident(true)
                            .withGroupingConfiguration(
                                new GroupingConfiguration()
                                    .withEnabled(true)
                                    .withReopenClosedIncident(false)
                                    .withLookbackDuration(Duration.parse("PT5H"))
                                    .withMatchingMethod(MatchingMethod.SELECTED)
                                    .withGroupByEntities(
                                        Arrays.asList(EntityMappingType.HOST, EntityMappingType.ACCOUNT))))
                    .withEventGroupingSettings(
                        new EventGroupingSettings().withAggregationKind(EventGroupingAggregationKind.ALERT_PER_RESULT)),
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
