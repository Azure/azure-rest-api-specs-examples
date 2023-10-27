import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.resources.fluent.models.PolicyAssignmentInner;
import com.azure.resourcemanager.resources.models.EnforcementMode;
import com.azure.resourcemanager.resources.models.Identity;
import com.azure.resourcemanager.resources.models.ParameterValuesValue;
import com.azure.resourcemanager.resources.models.ResourceIdentityType;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/** Samples for PolicyAssignments Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/createPolicyAssignmentWithIdentity.json
     */
    /**
     * Sample code: Create or update a policy assignment with a system assigned identity.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAPolicyAssignmentWithASystemAssignedIdentity(
        com.azure.resourcemanager.AzureResourceManager azure) throws IOException {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyAssignments()
            .createWithResponse(
                "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2",
                "EnforceNaming",
                new PolicyAssignmentInner()
                    .withLocation("eastus")
                    .withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
                    .withDisplayName("Enforce resource naming rules")
                    .withPolicyDefinitionId(
                        "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming")
                    .withParameters(
                        mapOf(
                            "prefix",
                            new ParameterValuesValue().withValue("DeptA"),
                            "suffix",
                            new ParameterValuesValue().withValue("-LC")))
                    .withDescription("Force resource names to begin with given DeptA and end with -LC")
                    .withMetadata(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize("{\"assignedBy\":\"Foo Bar\"}", Object.class, SerializerEncoding.JSON))
                    .withEnforcementMode(EnforcementMode.DEFAULT),
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
