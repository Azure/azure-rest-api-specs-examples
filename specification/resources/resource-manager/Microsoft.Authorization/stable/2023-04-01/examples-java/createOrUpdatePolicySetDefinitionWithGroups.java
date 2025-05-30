
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.resources.fluent.models.PolicySetDefinitionInner;
import com.azure.resourcemanager.resources.models.ParameterValuesValue;
import com.azure.resourcemanager.resources.models.PolicyDefinitionGroup;
import com.azure.resourcemanager.resources.models.PolicyDefinitionReference;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for PolicySetDefinitions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * createOrUpdatePolicySetDefinitionWithGroups.json
     */
    /**
     * Sample code: Create or update a policy set definition with groups.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAPolicySetDefinitionWithGroups(
        com.azure.resourcemanager.AzureResourceManager azure) throws IOException {
        azure.genericResources().manager().policyClient().getPolicySetDefinitions()
            .createOrUpdateWithResponse("CostManagement", new PolicySetDefinitionInner()
                .withDisplayName("Cost Management").withDescription("Policies to enforce low cost storage SKUs")
                .withMetadata(SerializerFactory.createDefaultManagementSerializerAdapter()
                    .deserialize("{\"category\":\"Cost Management\"}", Object.class, SerializerEncoding.JSON))
                .withPolicyDefinitions(Arrays.asList(new PolicyDefinitionReference().withPolicyDefinitionId(
                    "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1")
                    .withDefinitionVersion("1.*.*")
                    .withParameters(mapOf("listOfAllowedSKUs",
                        new ParameterValuesValue()
                            .withValue(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                                "[\"Standard_GRS\",\"Standard_LRS\"]", Object.class, SerializerEncoding.JSON))))
                    .withPolicyDefinitionReferenceId("Limit_Skus").withGroupNames(Arrays.asList("CostSaving")),
                    new PolicyDefinitionReference().withPolicyDefinitionId(
                        "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming")
                        .withDefinitionVersion("1.*.*")
                        .withParameters(mapOf("prefix", new ParameterValuesValue().withValue("DeptA"), "suffix",
                            new ParameterValuesValue().withValue("-LC")))
                        .withPolicyDefinitionReferenceId("Resource_Naming")
                        .withGroupNames(Arrays.asList("Organizational"))))
                .withPolicyDefinitionGroups(Arrays.asList(
                    new PolicyDefinitionGroup().withName("CostSaving").withDisplayName("Cost Management Policies")
                        .withDescription("Policies designed to control spend within a subscription."),
                    new PolicyDefinitionGroup().withName("Organizational").withDisplayName("Organizational Policies")
                        .withDescription(
                            "Policies that help enforce resource organization standards within a subscription."))),
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
