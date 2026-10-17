
import com.azure.resourcemanager.cognitiveservices.models.RaiRegoEncoding;
import com.azure.resourcemanager.cognitiveservices.models.RaiRegoProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for RaiRegos CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/PutRaiRego.json
     */
    /**
     * Sample code: Create a reusable Rego artifact.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        createAReusableRegoArtifact(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiRegos().define("input-guard").withExistingAccount("resource-group", "safety-account")
            .withTags(mapOf("environment", "production"))
            .withProperties(new RaiRegoProperties().withEncoding(RaiRegoEncoding.NONE).withRego(
                "package input_guard\n\nimport rego.v1\n\ndefault verdict := {\"decision\": \"allow\"}\n\nverdict := {\"decision\": \"deny\", \"reason\": \"prompt_injection_detected\"} if {\n  input.snapshot.moderation.harm.PromptInjection.detected\n}"))
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
