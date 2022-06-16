import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.resources.fluent.models.PolicyDefinitionInner;
import com.azure.resourcemanager.resources.models.ParameterDefinitionsValue;
import com.azure.resourcemanager.resources.models.ParameterDefinitionsValueMetadata;
import com.azure.resourcemanager.resources.models.ParameterType;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/** Samples for PolicyDefinitions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicyDefinition.json
     */
    /**
     * Sample code: Create or update a policy definition.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAPolicyDefinition(com.azure.resourcemanager.AzureResourceManager azure)
        throws IOException {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyDefinitions()
            .createOrUpdateWithResponse(
                "ResourceNaming",
                new PolicyDefinitionInner()
                    .withMode("All")
                    .withDisplayName("Enforce resource naming convention")
                    .withDescription("Force resource names to begin with given 'prefix' and/or end with given 'suffix'")
                    .withPolicyRule(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize(
                                "{\"if\":{\"not\":{\"field\":\"name\",\"like\":\"[concat(parameters('prefix'), '*',"
                                    + " parameters('suffix'))]\"}},\"then\":{\"effect\":\"deny\"}}",
                                Object.class,
                                SerializerEncoding.JSON))
                    .withMetadata(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize("{\"category\":\"Naming\"}", Object.class, SerializerEncoding.JSON))
                    .withParameters(
                        mapOf(
                            "prefix",
                            new ParameterDefinitionsValue()
                                .withType(ParameterType.STRING)
                                .withMetadata(
                                    new ParameterDefinitionsValueMetadata()
                                        .withDisplayName("Prefix")
                                        .withDescription("Resource name prefix")
                                        .withAdditionalProperties(mapOf())),
                            "suffix",
                            new ParameterDefinitionsValue()
                                .withType(ParameterType.STRING)
                                .withMetadata(
                                    new ParameterDefinitionsValueMetadata()
                                        .withDisplayName("Suffix")
                                        .withDescription("Resource name suffix")
                                        .withAdditionalProperties(mapOf())))),
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
