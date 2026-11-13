
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyProperties;

/**
 * Samples for RaiPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/PutRaiPolicyAcsWithoutTools.json
     */
    /**
     * Sample code: Create an ACS policy without a tool catalog.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createAnACSPolicyWithoutAToolCatalog(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiPolicies().define("input-guard").withExistingAccount("resource-group", "safety-account")
            .withProperties(new RaiPolicyProperties()).create();
    }
}
