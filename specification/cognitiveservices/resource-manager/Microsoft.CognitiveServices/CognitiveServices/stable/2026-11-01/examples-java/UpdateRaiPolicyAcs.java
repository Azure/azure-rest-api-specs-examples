
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyProperties;
import java.util.Arrays;

/**
 * Samples for RaiPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/UpdateRaiPolicyAcs.json
     */
    /**
     * Sample code: Replace an ACS policy conditionally.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        replaceAnACSPolicyConditionally(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiPolicies().define("agent-guard").withExistingAccount("resource-group", "safety-account")
            .withProperties(new RaiPolicyProperties().withCustomBlocklists(Arrays.asList())).create();
    }
}
