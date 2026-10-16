
import com.azure.resourcemanager.cognitiveservices.models.ContentLevel;
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyContentFilter;
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyContentSource;
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyMode;
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for RaiPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/PutRaiPolicy.json
     */
    /**
     * Sample code: PutRaiPolicy.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void putRaiPolicy(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiPolicies().define("raiPolicyName").withExistingAccount("resourceGroupName", "accountName")
            .withProperties(new RaiPolicyProperties().withMode(RaiPolicyMode.ASYNCHRONOUS_FILTER)
                .withBasePolicyName("Microsoft.Default")
                .withContentFilters(Arrays.asList(
                    new RaiPolicyContentFilter().withName("Hate").withEnabled(false)
                        .withSeverityThreshold(ContentLevel.HIGH).withBlocking(false)
                        .withSource(RaiPolicyContentSource.PROMPT),
                    new RaiPolicyContentFilter().withName("Hate").withEnabled(true)
                        .withSeverityThreshold(ContentLevel.MEDIUM).withBlocking(true)
                        .withSource(RaiPolicyContentSource.COMPLETION),
                    new RaiPolicyContentFilter().withName("Sexual").withEnabled(true)
                        .withSeverityThreshold(ContentLevel.HIGH).withBlocking(true)
                        .withSource(RaiPolicyContentSource.PROMPT),
                    new RaiPolicyContentFilter().withName("Sexual").withEnabled(true)
                        .withSeverityThreshold(ContentLevel.MEDIUM).withBlocking(true)
                        .withSource(RaiPolicyContentSource.COMPLETION),
                    new RaiPolicyContentFilter().withName("Selfharm").withEnabled(true)
                        .withSeverityThreshold(ContentLevel.HIGH).withBlocking(true)
                        .withSource(RaiPolicyContentSource.PROMPT),
                    new RaiPolicyContentFilter().withName("Selfharm").withEnabled(true)
                        .withSeverityThreshold(ContentLevel.MEDIUM).withBlocking(true)
                        .withSource(RaiPolicyContentSource.COMPLETION),
                    new RaiPolicyContentFilter().withName("Violence").withEnabled(true)
                        .withSeverityThreshold(ContentLevel.MEDIUM).withBlocking(true)
                        .withSource(RaiPolicyContentSource.PROMPT),
                    new RaiPolicyContentFilter().withName("Violence").withEnabled(true)
                        .withSeverityThreshold(ContentLevel.MEDIUM).withBlocking(true)
                        .withSource(RaiPolicyContentSource.COMPLETION),
                    new RaiPolicyContentFilter().withName("Jailbreak").withEnabled(true).withBlocking(true)
                        .withSource(RaiPolicyContentSource.PROMPT),
                    new RaiPolicyContentFilter().withName("Protected Material Text").withEnabled(true)
                        .withBlocking(true).withSource(RaiPolicyContentSource.COMPLETION),
                    new RaiPolicyContentFilter().withName("Protected Material Code").withEnabled(true)
                        .withBlocking(true).withSource(RaiPolicyContentSource.COMPLETION),
                    new RaiPolicyContentFilter().withName("Profanity").withEnabled(true).withBlocking(true)
                        .withSource(RaiPolicyContentSource.PROMPT))))
            .create();
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
