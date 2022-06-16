import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.MicrosoftSecurityIncidentCreationAlertRule;
import com.azure.resourcemanager.securityinsights.models.MicrosoftSecurityProductName;
import java.util.HashMap;
import java.util.Map;

/** Samples for AlertRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/alertRules/CreateMicrosoftSecurityIncidentCreationAlertRule.json
     */
    /**
     * Sample code: Creates or updates a MicrosoftSecurityIncidentCreation rule.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAMicrosoftSecurityIncidentCreationRule(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .alertRules()
            .createOrUpdateWithResponse(
                "myRg",
                "myWorkspace",
                "microsoftSecurityIncidentCreationRuleExample",
                new MicrosoftSecurityIncidentCreationAlertRule()
                    .withEtag("\"260097e0-0000-0d00-0000-5d6fa88f0000\"")
                    .withDisplayName("testing displayname")
                    .withEnabled(true)
                    .withProductFilter(MicrosoftSecurityProductName.MICROSOFT_CLOUD_APP_SECURITY),
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
