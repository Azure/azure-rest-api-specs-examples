
import com.azure.resourcemanager.alertsmanagement.models.AddActionGroups;
import com.azure.resourcemanager.alertsmanagement.models.AlertProcessingRuleProperties;
import com.azure.resourcemanager.alertsmanagement.models.Condition;
import com.azure.resourcemanager.alertsmanagement.models.Field;
import com.azure.resourcemanager.alertsmanagement.models.Operator;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AlertProcessingRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/
     * AlertProcessingRules_Create_or_update_add_two_action_groups_all_Sev0_Sev1_two_resource_groups.json
     */
    /**
     * Sample code: Create or update a rule that adds two action groups to all Sev0 and Sev1 alerts in two resource
     * groups.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void createOrUpdateARuleThatAddsTwoActionGroupsToAllSev0AndSev1AlertsInTwoResourceGroups(
        com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alertProcessingRules().define("AddActionGroupsBySeverity").withRegion("Global")
            .withExistingResourceGroup("alertscorrelationrg").withTags(mapOf())
            .withProperties(new AlertProcessingRuleProperties()
                .withScopes(Arrays.asList("/subscriptions/subId1/resourceGroups/RGId1",
                    "/subscriptions/subId1/resourceGroups/RGId2"))
                .withConditions(Arrays.asList(new Condition().withField(Field.SEVERITY).withOperator(Operator.EQUALS)
                    .withValues(Arrays.asList("sev0", "sev1"))))
                .withActions(Arrays.asList(new AddActionGroups().withActionGroupIds(Arrays.asList(
                    "/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/AGId1",
                    "/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/AGId2"))))
                .withDescription("Add AGId1 and AGId2 to all Sev0 and Sev1 alerts in these resourceGroups")
                .withEnabled(true))
            .create();
    }

    // Use "Map.of" if available
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
