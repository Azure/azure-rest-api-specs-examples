
import com.azure.resourcemanager.resources.models.OverrideKind;
import com.azure.resourcemanager.resources.models.OverrideModel;
import com.azure.resourcemanager.resources.models.PolicyAssignmentUpdate;
import com.azure.resourcemanager.resources.models.Selector;
import com.azure.resourcemanager.resources.models.SelectorKind;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for PolicyAssignments Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * updatePolicyAssignmentWithOverrides.json
     */
    /**
     * Sample code: Update a policy assignment with overrides.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAPolicyAssignmentWithOverrides(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyAssignments().updateWithResponse(
            "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2", "CostManagement",
            new PolicyAssignmentUpdate()
                .withOverrides(Arrays.asList(new OverrideModel().withKind(OverrideKind.POLICY_EFFECT).withValue("Audit")
                    .withSelectors(Arrays.asList(new Selector().withKind(SelectorKind.POLICY_DEFINITION_REFERENCE_ID)
                        .withIn(Arrays.asList("Limit_Skus", "Limit_Locations")))))),
            com.azure.core.util.Context.NONE);
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
