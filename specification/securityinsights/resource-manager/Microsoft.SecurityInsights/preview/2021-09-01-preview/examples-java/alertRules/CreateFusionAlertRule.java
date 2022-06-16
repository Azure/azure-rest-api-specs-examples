import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.FusionAlertRule;
import java.util.HashMap;
import java.util.Map;

/** Samples for AlertRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/alertRules/CreateFusionAlertRule.json
     */
    /**
     * Sample code: Creates or updates a Fusion alert rule.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAFusionAlertRule(
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
                    .withEnabled(true),
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
