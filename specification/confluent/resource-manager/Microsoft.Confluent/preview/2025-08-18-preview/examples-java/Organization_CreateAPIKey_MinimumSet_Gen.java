
import com.azure.resourcemanager.confluent.models.CreateApiKeyModel;

/**
 * Samples for Organization CreateApiKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Organization_CreateAPIKey_MinimumSet_Gen.json
     */
    /**
     * Sample code: Organization_CreateAPIKey_MinimumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationCreateAPIKeyMinimumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().createApiKeyWithResponse("rgconfluent", "qbnpbkqxwtvjnytnconwynln", "un",
            "vuwuoryynnsuyfkicyejllc", new CreateApiKeyModel(), com.azure.core.util.Context.NONE);
    }
}
