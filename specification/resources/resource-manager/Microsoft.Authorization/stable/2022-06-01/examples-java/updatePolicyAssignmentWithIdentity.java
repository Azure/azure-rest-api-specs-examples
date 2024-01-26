
import com.azure.resourcemanager.resources.models.Identity;
import com.azure.resourcemanager.resources.models.PolicyAssignmentUpdate;
import com.azure.resourcemanager.resources.models.ResourceIdentityType;
import java.util.HashMap;
import java.util.Map;

/** Samples for PolicyAssignments Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/
     * updatePolicyAssignmentWithIdentity.json
     */
    /**
     * Sample code: Update a policy assignment with a system assigned identity.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        updateAPolicyAssignmentWithASystemAssignedIdentity(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyAssignments().updateWithResponse(
            "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2", "EnforceNaming", new PolicyAssignmentUpdate()
                .withLocation("eastus").withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED)),
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
