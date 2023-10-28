import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.resources.fluent.models.PolicyAssignmentInner;
import com.azure.resourcemanager.resources.models.ResourceSelector;
import com.azure.resourcemanager.resources.models.Selector;
import com.azure.resourcemanager.resources.models.SelectorKind;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for PolicyAssignments Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/createPolicyAssignmentWithResourceSelectors.json
     */
    /**
     * Sample code: Create or update a policy assignment with resource selectors.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAPolicyAssignmentWithResourceSelectors(
        com.azure.resourcemanager.AzureResourceManager azure) throws IOException {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyAssignments()
            .createWithResponse(
                "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2",
                "CostManagement",
                new PolicyAssignmentInner()
                    .withDisplayName("Limit the resource location and resource SKU")
                    .withPolicyDefinitionId(
                        "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/CostManagement")
                    .withDescription("Limit the resource location and resource SKU")
                    .withMetadata(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize("{\"assignedBy\":\"Special Someone\"}", Object.class, SerializerEncoding.JSON))
                    .withResourceSelectors(
                        Arrays
                            .asList(
                                new ResourceSelector()
                                    .withName("SDPRegions")
                                    .withSelectors(
                                        Arrays
                                            .asList(
                                                new Selector()
                                                    .withKind(SelectorKind.RESOURCE_LOCATION)
                                                    .withIn(Arrays.asList("eastus2euap", "centraluseuap")))))),
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
