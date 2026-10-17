
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsInterventionPoint;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsInterventionPoints;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsManifest;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsPolicyBinding;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsPolicyDefinitionType;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsPolicyTarget;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsPolicyTargetKind;
import com.azure.resourcemanager.cognitiveservices.models.RaiAcsRegoPolicyDefinition;
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyFormat;
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
     * x-ms-original-file: 2026-09-15-preview/PutRaiPolicyAcsWithoutTools.json
     */
    /**
     * Sample code: Create an ACS policy without a tool catalog.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createAnACSPolicyWithoutAToolCatalog(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiPolicies().define("input-guard").withExistingAccount("resource-group", "safety-account")
            .withProperties(new RaiPolicyProperties().withFormat(RaiPolicyFormat.ACS).withAcs(new RaiAcsManifest()
                .withAgentControlSpecificationVersion("0.4.0-alpha.1")
                .withPolicies(mapOf("input-guard",
                    new RaiAcsRegoPolicyDefinition().withType(RaiAcsPolicyDefinitionType.REGO)
                        .withQuery("data.input_guard.verdict")))
                .withInterventionPoints(new RaiAcsInterventionPoints().withInput(new RaiAcsInterventionPoint()
                    .withPolicyTarget(RaiAcsPolicyTarget.INPUT).withPolicyTargetKind(RaiAcsPolicyTargetKind.USER_INPUT)
                    .withPolicy(new RaiAcsPolicyBinding().withId("input-guard")))))
                .withAcsRegos(Arrays.asList(new RaiRegoReference().withRegoName("input-guard"))))
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
