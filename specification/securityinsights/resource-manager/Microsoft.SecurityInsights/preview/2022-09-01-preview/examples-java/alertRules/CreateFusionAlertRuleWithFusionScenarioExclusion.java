import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.AlertSeverity;
import com.azure.resourcemanager.securityinsights.models.FusionAlertRule;
import com.azure.resourcemanager.securityinsights.models.FusionSourceSettings;
import com.azure.resourcemanager.securityinsights.models.FusionSourceSubTypeSetting;
import com.azure.resourcemanager.securityinsights.models.FusionSubTypeSeverityFilter;
import com.azure.resourcemanager.securityinsights.models.FusionSubTypeSeverityFiltersItem;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for AlertRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/CreateFusionAlertRuleWithFusionScenarioExclusion.json
     */
    /**
     * Sample code: Creates or updates a Fusion alert rule with scenario exclusion pattern.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAFusionAlertRuleWithScenarioExclusionPattern(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .alertRules()
            .createOrUpdateWithResponse(
                "myRg",
                "myWorkspace",
                "myFirstFusionRule",
                new FusionAlertRule()
                    .withEtag("3d00c3ca-0000-0100-0000-5d42d5010000")
                    .withAlertRuleTemplateName("f71aba3d-28fb-450b-b192-4e76a83015c8")
                    .withEnabled(true)
                    .withSourceSettings(
                        Arrays
                            .asList(
                                new FusionSourceSettings().withEnabled(true).withSourceName("Anomalies"),
                                new FusionSourceSettings()
                                    .withEnabled(true)
                                    .withSourceName("Alert providers")
                                    .withSourceSubTypes(
                                        Arrays
                                            .asList(
                                                new FusionSourceSubTypeSetting()
                                                    .withEnabled(true)
                                                    .withSourceSubTypeName("Azure Active Directory Identity Protection")
                                                    .withSeverityFilters(
                                                        new FusionSubTypeSeverityFilter()
                                                            .withFilters(
                                                                Arrays
                                                                    .asList(
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.HIGH)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.MEDIUM)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.LOW)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.INFORMATIONAL)
                                                                            .withEnabled(true)))),
                                                new FusionSourceSubTypeSetting()
                                                    .withEnabled(true)
                                                    .withSourceSubTypeName("Azure Defender")
                                                    .withSeverityFilters(
                                                        new FusionSubTypeSeverityFilter()
                                                            .withFilters(
                                                                Arrays
                                                                    .asList(
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.HIGH)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.MEDIUM)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.LOW)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.INFORMATIONAL)
                                                                            .withEnabled(true)))),
                                                new FusionSourceSubTypeSetting()
                                                    .withEnabled(true)
                                                    .withSourceSubTypeName("Azure Defender for IoT")
                                                    .withSeverityFilters(
                                                        new FusionSubTypeSeverityFilter()
                                                            .withFilters(
                                                                Arrays
                                                                    .asList(
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.HIGH)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.MEDIUM)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.LOW)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.INFORMATIONAL)
                                                                            .withEnabled(true)))),
                                                new FusionSourceSubTypeSetting()
                                                    .withEnabled(true)
                                                    .withSourceSubTypeName("Microsoft 365 Defender")
                                                    .withSeverityFilters(
                                                        new FusionSubTypeSeverityFilter()
                                                            .withFilters(
                                                                Arrays
                                                                    .asList(
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.HIGH)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.MEDIUM)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.LOW)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.INFORMATIONAL)
                                                                            .withEnabled(true)))),
                                                new FusionSourceSubTypeSetting()
                                                    .withEnabled(true)
                                                    .withSourceSubTypeName("Microsoft Cloud App Security")
                                                    .withSeverityFilters(
                                                        new FusionSubTypeSeverityFilter()
                                                            .withFilters(
                                                                Arrays
                                                                    .asList(
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.HIGH)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.MEDIUM)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.LOW)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.INFORMATIONAL)
                                                                            .withEnabled(true)))),
                                                new FusionSourceSubTypeSetting()
                                                    .withEnabled(true)
                                                    .withSourceSubTypeName("Microsoft Defender for Endpoint")
                                                    .withSeverityFilters(
                                                        new FusionSubTypeSeverityFilter()
                                                            .withFilters(
                                                                Arrays
                                                                    .asList(
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.HIGH)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.MEDIUM)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.LOW)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.INFORMATIONAL)
                                                                            .withEnabled(true)))),
                                                new FusionSourceSubTypeSetting()
                                                    .withEnabled(true)
                                                    .withSourceSubTypeName("Microsoft Defender for Identity")
                                                    .withSeverityFilters(
                                                        new FusionSubTypeSeverityFilter()
                                                            .withFilters(
                                                                Arrays
                                                                    .asList(
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.HIGH)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.MEDIUM)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.LOW)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.INFORMATIONAL)
                                                                            .withEnabled(true)))),
                                                new FusionSourceSubTypeSetting()
                                                    .withEnabled(true)
                                                    .withSourceSubTypeName("Microsoft Defender for Office 365")
                                                    .withSeverityFilters(
                                                        new FusionSubTypeSeverityFilter()
                                                            .withFilters(
                                                                Arrays
                                                                    .asList(
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.HIGH)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.MEDIUM)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.LOW)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.INFORMATIONAL)
                                                                            .withEnabled(true)))),
                                                new FusionSourceSubTypeSetting()
                                                    .withEnabled(true)
                                                    .withSourceSubTypeName("Azure Sentinel scheduled analytics rules")
                                                    .withSeverityFilters(
                                                        new FusionSubTypeSeverityFilter()
                                                            .withFilters(
                                                                Arrays
                                                                    .asList(
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.HIGH)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.MEDIUM)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.LOW)
                                                                            .withEnabled(true),
                                                                        new FusionSubTypeSeverityFiltersItem()
                                                                            .withSeverity(AlertSeverity.INFORMATIONAL)
                                                                            .withEnabled(true)))))),
                                new FusionSourceSettings()
                                    .withEnabled(true)
                                    .withSourceName("Raw logs from other sources")
                                    .withSourceSubTypes(
                                        Arrays
                                            .asList(
                                                new FusionSourceSubTypeSetting()
                                                    .withEnabled(true)
                                                    .withSourceSubTypeName("Palo Alto Networks")
                                                    .withSeverityFilters(new FusionSubTypeSeverityFilter()))))),
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
