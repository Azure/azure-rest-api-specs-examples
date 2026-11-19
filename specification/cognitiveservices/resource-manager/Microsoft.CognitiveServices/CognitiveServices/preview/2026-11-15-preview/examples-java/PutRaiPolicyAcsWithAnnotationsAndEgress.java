
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsAgentShutdownInterventionPoint;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsAgentStartupInterventionPoint;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsAnnotationBinding;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsAnnotatorDefinition;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsAnnotatorType;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsInterventionPoint;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsInterventionPoints;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsManifest;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsPolicyBinding;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsPolicyDefinitionType;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsPolicyTarget;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsPolicyTargetKind;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsPostModelCallInterventionPoint;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsPreModelCallInterventionPoint;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsRegoPolicyDefinition;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsToolInterventionPoint;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressDefaultAction;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressHeaderOperation;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressHeaderTransform;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressMode;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressPolicyConfig;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressRule;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressRuleAction;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressRuleActionType;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressRuleMatch;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressRuleType;
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyFormat;
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyMode;
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyProperties;
import com.azure.resourcemanager.cognitiveservices.models.RaiRegoReference;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for RaiPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/PutRaiPolicyAcsWithAnnotationsAndEgress.json
     */
    /**
     * Sample code: Create an ACS policy with optional annotations and an egress companion.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createAnACSPolicyWithOptionalAnnotationsAndAnEgressCompanion(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiPolicies().define("annotated-egress-guard").withExistingAccount("resource-group", "safety-account")
            .withProperties(new RaiPolicyProperties().withFormat(RaiPolicyFormat.ACS).withAcs(new RaiAcsManifest()
                .withAgentControlSpecificationVersion("0.4.0-alpha.1")
                .withMetadata(mapOf("description",
                    "Illustrative annotation authoring; supported annotators and configuration depend on the evaluation host."))
                .withPolicies(mapOf("agent-guard",
                    new RaiAcsRegoPolicyDefinition().withType(RaiAcsPolicyDefinitionType.REGO)
                        .withQuery("data.agent_guard.verdict")))
                .withInterventionPoints(new RaiAcsInterventionPoints()
                    .withAgentStartup(new RaiAcsAgentStartupInterventionPoint()
                        .withPolicy(new RaiAcsPolicyBinding().withId("agent-guard")))
                    .withInput(new RaiAcsInterventionPoint().withPolicyTarget(RaiAcsPolicyTarget.INPUT)
                        .withPolicyTargetKind(RaiAcsPolicyTargetKind.USER_INPUT)
                        .withPolicy(new RaiAcsPolicyBinding().withId("agent-guard"))
                        .withAnnotations(mapOf("risk",
                            new RaiAcsAnnotationBinding().withFrom("$target.content")
                                .withAdditionalProperties(mapOf()))))
                    .withPreModelCall(new RaiAcsPreModelCallInterventionPoint()
                        .withPolicy(new RaiAcsPolicyBinding().withId("agent-guard")))
                    .withPostModelCall(new RaiAcsPostModelCallInterventionPoint()
                        .withPolicy(new RaiAcsPolicyBinding().withId("agent-guard")))
                    .withPreToolCall(
                        new RaiAcsToolInterventionPoint().withPolicyTarget(RaiAcsPolicyTarget.TOOL_ARGUMENTS)
                            .withPolicyTargetKind(RaiAcsPolicyTargetKind.TOOL_ARGUMENTS)
                            .withPolicy(new RaiAcsPolicyBinding().withId("agent-guard")))
                    .withPostToolCall(new RaiAcsToolInterventionPoint().withPolicyTarget(RaiAcsPolicyTarget.TOOL_RESULT)
                        .withPolicyTargetKind(RaiAcsPolicyTargetKind.TOOL_RESULT)
                        .withPolicy(new RaiAcsPolicyBinding().withId("agent-guard")))
                    .withOutput(new RaiAcsInterventionPoint().withPolicyTarget(RaiAcsPolicyTarget.OUTPUT)
                        .withPolicyTargetKind(RaiAcsPolicyTargetKind.ASSISTANT_OUTPUT)
                        .withPolicy(new RaiAcsPolicyBinding().withId("agent-guard"))
                        .withAnnotations(mapOf("risk",
                            new RaiAcsAnnotationBinding().withFrom("$target.content")
                                .withAdditionalProperties(mapOf()))))
                    .withAgentShutdown(new RaiAcsAgentShutdownInterventionPoint()
                        .withPolicy(new RaiAcsPolicyBinding().withId("agent-guard"))))
                .withAnnotators(mapOf("risk", new RaiAcsAnnotatorDefinition().withType(RaiAcsAnnotatorType.LLM)
                    .withAdditionalProperties(mapOf("endpoint", "https://synthetic-judge.openai.azure.com/", "provider",
                        "azure_openai", "result_schema", "acs.judge.observation/v2", "system_prompt",
                        "Classify the complete subject for harmful intent. Score from 0 to 10: 0 is benign, 5 is ambiguous or potentially harmful, and 10 is clear harmful intent. Use label clear for scores below 4, review for scores from 4 to below 7, and high_risk for scores of 7 or higher. Report observations only; Rego decides the action.",
                        "timeout_ms", 10000, "api_version", "2024-10-21", "deployment", "judge")))))
                .withAcsRegos(Arrays.asList(new RaiRegoReference().withRegoName("agent-guard")))
                .withMode(RaiPolicyMode.ASYNCHRONOUS_FILTER).withEgressPolicy(
                    new RaiEgressPolicyConfig().withMode(RaiEgressMode.AUDIT)
                        .withDefaultAction(RaiEgressDefaultAction.DENY)
                        .withRules(Arrays.asList(
                            new RaiEgressRule().withName("allow-docs").withRuleType(RaiEgressRuleType.FQDN)
                                .withMatch(
                                    new RaiEgressRuleMatch().withHost("learn.microsoft.com").withPath("/azure/*"))
                                .withAction(new RaiEgressRuleAction().withActionType(RaiEgressRuleActionType.ALLOW)),
                            new RaiEgressRule().withName("label-api").withRuleType(RaiEgressRuleType.FQDN)
                                .withMatch(new RaiEgressRuleMatch().withHost("api.contoso.com").withPath("/v1/*"))
                                .withAction(new RaiEgressRuleAction().withActionType(RaiEgressRuleActionType.TRANSFORM)
                                    .withHeaders(Arrays.asList(
                                        new RaiEgressHeaderTransform().withOperation(RaiEgressHeaderOperation.SET)
                                            .withName("X-Workload-Class").withValue("public-demo"))))))))
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
