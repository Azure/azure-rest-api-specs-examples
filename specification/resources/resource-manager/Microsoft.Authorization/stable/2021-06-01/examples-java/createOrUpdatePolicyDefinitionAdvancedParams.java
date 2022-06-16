import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.resources.fluent.models.PolicyDefinitionInner;
import com.azure.resourcemanager.resources.models.ParameterDefinitionsValue;
import com.azure.resourcemanager.resources.models.ParameterDefinitionsValueMetadata;
import com.azure.resourcemanager.resources.models.ParameterType;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for PolicyDefinitions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicyDefinitionAdvancedParams.json
     */
    /**
     * Sample code: Create or update a policy definition with advanced parameters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAPolicyDefinitionWithAdvancedParameters(
        com.azure.resourcemanager.AzureResourceManager azure) throws IOException {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyDefinitions()
            .createOrUpdateWithResponse(
                "EventHubDiagnosticLogs",
                new PolicyDefinitionInner()
                    .withMode("Indexed")
                    .withDisplayName("Event Hubs should have diagnostic logging enabled")
                    .withDescription(
                        "Audit enabling of logs and retain them up to a year. This enables recreation of activity"
                            + " trails for investigation purposes when a security incident occurs or your network is"
                            + " compromised")
                    .withPolicyRule(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize(
                                "{\"if\":{\"equals\":\"Microsoft.EventHub/namespaces\",\"field\":\"type\"},\"then\":{\"effect\":\"AuditIfNotExists\",\"details\":{\"type\":\"Microsoft.Insights/diagnosticSettings\",\"existenceCondition\":{\"allOf\":[{\"equals\":\"true\",\"field\":\"Microsoft.Insights/diagnosticSettings/logs[*].retentionPolicy.enabled\"},{\"equals\":\"[parameters('requiredRetentionDays')]\",\"field\":\"Microsoft.Insights/diagnosticSettings/logs[*].retentionPolicy.days\"}]}}}}",
                                Object.class,
                                SerializerEncoding.JSON))
                    .withMetadata(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize("{\"category\":\"Event Hub\"}", Object.class, SerializerEncoding.JSON))
                    .withParameters(
                        mapOf(
                            "requiredRetentionDays",
                            new ParameterDefinitionsValue()
                                .withType(ParameterType.INTEGER)
                                .withAllowedValues(Arrays.asList(0, 30, 90, 180, 365))
                                .withDefaultValue(365)
                                .withMetadata(
                                    new ParameterDefinitionsValueMetadata()
                                        .withDisplayName("Required retention (days)")
                                        .withDescription("The required diagnostic logs retention in days")
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
