Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.resources.fluent.models.PolicyAssignmentInner;
import com.azure.resourcemanager.resources.models.NonComplianceMessage;
import com.azure.resourcemanager.resources.models.ParameterValuesValue;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for PolicyAssignments Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createPolicyAssignment.json
     */
    /**
     * Sample code: Create or update a policy assignment.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAPolicyAssignment(com.azure.resourcemanager.AzureResourceManager azure)
        throws IOException {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyAssignments()
            .createWithResponse(
                "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2",
                "EnforceNaming",
                new PolicyAssignmentInner()
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
                            .deserialize("{\"assignedBy\":\"Special Someone\"}", Object.class, SerializerEncoding.JSON))
                    .withNonComplianceMessages(
                        Arrays
                            .asList(
                                new NonComplianceMessage()
                                    .withMessage("Resource names must start with 'DeptA' and end with '-LC'."))),
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
```
