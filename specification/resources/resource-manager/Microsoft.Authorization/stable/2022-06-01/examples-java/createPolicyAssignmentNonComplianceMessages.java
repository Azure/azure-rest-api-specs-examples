
import com.azure.resourcemanager.resources.fluent.models.PolicyAssignmentInner;
import com.azure.resourcemanager.resources.models.NonComplianceMessage;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for PolicyAssignments Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/
     * createPolicyAssignmentNonComplianceMessages.json
     */
    /**
     * Sample code: Create or update a policy assignment with multiple non-compliance messages.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAPolicyAssignmentWithMultipleNonComplianceMessages(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyAssignments().createWithResponse(
            "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2", "securityInitAssignment",
            new PolicyAssignmentInner().withDisplayName("Enforce security policies").withPolicyDefinitionId(
                "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/securityInitiative")
                .withNonComplianceMessages(Arrays.asList(
                    new NonComplianceMessage()
                        .withMessage("Resources must comply with all internal security policies. See <internal site"
                            + " URL> for more info."),
                    new NonComplianceMessage().withMessage("Resource names must start with 'DeptA' and end with '-LC'.")
                        .withPolicyDefinitionReferenceId("10420126870854049575"),
                    new NonComplianceMessage().withMessage("Storage accounts must have firewall rules configured.")
                        .withPolicyDefinitionReferenceId("8572513655450389710"))),
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
