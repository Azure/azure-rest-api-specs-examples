
import com.azure.resourcemanager.cognitiveservices.models.CustomBlocklistConfig;
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyContentSource;
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyProperties;
import java.util.Arrays;

/**
 * Samples for RaiPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/PutRaiPolicyAcs.json
     */
    /**
     * Sample code: Create an ACS policy with optional policy dependencies.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createAnACSPolicyWithOptionalPolicyDependencies(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiPolicies().define("agent-guard").withExistingAccount("resource-group", "safety-account")
            .withProperties(new RaiPolicyProperties().withCustomBlocklists(Arrays.asList(new CustomBlocklistConfig()
                .withBlocklistName("blocked-terms").withBlocking(true).withSource(RaiPolicyContentSource.PROMPT))))
            .create();
    }
}
