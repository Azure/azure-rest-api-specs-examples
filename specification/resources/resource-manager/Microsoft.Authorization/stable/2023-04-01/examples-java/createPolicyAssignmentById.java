
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.resources.fluent.models.PolicyAssignmentInner;
import com.azure.resourcemanager.resources.models.EnforcementMode;
import com.azure.resourcemanager.resources.models.ParameterValuesValue;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for PolicyAssignments CreateById.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * createPolicyAssignmentById.json
     */
    /**
     * Sample code: Create or update policy assignment by ID.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdatePolicyAssignmentByID(com.azure.resourcemanager.AzureResourceManager azure)
        throws IOException {
        azure.genericResources().manager().policyClient().getPolicyAssignments().createByIdWithResponse(
            "providers/Microsoft.Management/managementGroups/MyManagementGroup/providers/Microsoft.Authorization/policyAssignments/LowCostStorage",
            new PolicyAssignmentInner().withDisplayName("Enforce storage account SKU")
                .withPolicyDefinitionId(
                    "/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1")
                .withDefinitionVersion("1.*.*")
                .withParameters(mapOf("listOfAllowedSKUs",
                    new ParameterValuesValue().withValue(SerializerFactory.createDefaultManagementSerializerAdapter()
                        .deserialize("[\"Standard_GRS\",\"Standard_LRS\"]", Object.class, SerializerEncoding.JSON))))
                .withDescription("Allow only storage accounts of SKU Standard_GRS or Standard_LRS to be created")
                .withMetadata(SerializerFactory.createDefaultManagementSerializerAdapter()
                    .deserialize("{\"assignedBy\":\"Cheapskate Boss\"}", Object.class, SerializerEncoding.JSON))
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
