
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.cognitiveservices.models.CustomBlocklistConfig;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsEmptyObject;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsHarmCategory;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsHarmConfiguration;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsInterventionPoint;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsInterventionPoints;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsManifest;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsModerationBindingExtension;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsModerationSubjectFormat;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsPolicyBinding;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsPolicyDefinitionType;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsPolicyTarget;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsPolicyTargetKind;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsRegoPolicyDefinition;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsToolDefinition;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsToolInterventionPoint;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsToolNameSelector;
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyContentSource;
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyCustomExternalSafetyProviderReference;
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyFormat;
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyProperties;
import com.azure.resourcemanager.cognitiveservices.models.RaiRegoReference;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for RaiPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/PutRaiPolicyAcs.json
     */
    /**
     * Sample code: Create an ACS policy with optional policy dependencies.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createAnACSPolicyWithOptionalPolicyDependencies(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) throws IOException {
        manager.raiPolicies().define("agent-guard").withExistingAccount("resource-group", "safety-account")
            .withProperties(new RaiPolicyProperties().withFormat(RaiPolicyFormat.ACS)
                .withAcs(new RaiAcsManifest().withAgentControlSpecificationVersion("0.4.0-alpha.1")
                    .withMetadata(mapOf("name", "agent-guard"))
                    .withPolicies(mapOf("input-guard",
                        new RaiAcsRegoPolicyDefinition()
                            .withType(RaiAcsPolicyDefinitionType.REGO).withQuery("data.input_guard.verdict"),
                        "tool-guard",
                        new RaiAcsRegoPolicyDefinition().withType(RaiAcsPolicyDefinitionType.REGO)
                            .withQuery("data.tool_guard.verdict")))
                    .withInterventionPoints(new RaiAcsInterventionPoints()
                        .withInput(new RaiAcsInterventionPoint().withPolicyTarget(RaiAcsPolicyTarget.INPUT)
                            .withPolicyTargetKind(RaiAcsPolicyTargetKind.USER_INPUT)
                            .withPolicy(new RaiAcsPolicyBinding().withId("input-guard")
                                .withAacsModeration(new RaiAcsModerationBindingExtension()
                                    .withSubjectFormat(RaiAcsModerationSubjectFormat.TEXT)
                                    .withHarmConfigs(Arrays.asList(new RaiAcsHarmConfiguration()
                                        .withCategory(RaiAcsHarmCategory.PROMPT_INJECTION))))))
                        .withPreToolCall(
                            new RaiAcsToolInterventionPoint().withPolicyTarget(RaiAcsPolicyTarget.TOOL_ARGUMENTS)
                                .withPolicyTargetKind(RaiAcsPolicyTargetKind.TOOL_ARGUMENTS)
                                .withPolicy(new RaiAcsPolicyBinding().withId("tool-guard"))
                                .withToolNameFrom(RaiAcsToolNameSelector.TOOL_CALL_NAME))
                        .withPostToolCall(
                            new RaiAcsToolInterventionPoint().withPolicyTarget(RaiAcsPolicyTarget.TOOL_RESULT)
                                .withPolicyTargetKind(RaiAcsPolicyTargetKind.TOOL_RESULT)
                                .withPolicy(new RaiAcsPolicyBinding().withId("tool-guard"))
                                .withToolNameFrom(RaiAcsToolNameSelector.TOOL_CALL_NAME)))
                    .withTools(mapOf("web_search",
                        new RaiAcsToolDefinition().withId("web_search").withType("retrieval")
                            .withDescription("Search approved public documentation")
                            .withSecurityLabels(Arrays.asList("network_egress", "untrusted_content"))
                            .withClearance("public")
                            .withAdditionalProperties(mapOf("allowed_domains",
                                SerializerFactory.createDefaultManagementSerializerAdapter()
                                    .deserialize("[\"learn.microsoft.com\"]", Object.class, SerializerEncoding.JSON))),
                        "wire_transfer",
                        new RaiAcsToolDefinition().withId("wire_transfer").withType("financial_action")
                            .withSecurityLabels(Arrays.asList("financial_write")).withClearance("confidential")
                            .withAdditionalProperties(mapOf())))
                    .withAnnotators(new RaiAcsEmptyObject()))
                .withAcsRegos(Arrays.asList(new RaiRegoReference().withRegoName("input-guard"),
                    new RaiRegoReference().withRegoName("tool-guard")))
                .withCustomBlocklists(Arrays.asList(new CustomBlocklistConfig().withBlocklistName("blocked-terms")
                    .withBlocking(true).withSource(RaiPolicyContentSource.PROMPT)))
                .withCustomExternalSafetyProviders(Arrays.asList(new RaiPolicyCustomExternalSafetyProviderReference()
                    .withExternalSafetyProviderName("contoso-safety-provider")
                    .withManagedIdentityResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/safety-provider-identity")
                    .withSource(RaiPolicyContentSource.PROMPT).withBlocking(true))))
            .withIfNoneMatch("*").create();
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
