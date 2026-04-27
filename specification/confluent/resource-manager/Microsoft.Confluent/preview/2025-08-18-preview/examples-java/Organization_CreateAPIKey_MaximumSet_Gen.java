
import com.azure.resourcemanager.confluent.models.CreateApiKeyModel;

/**
 * Samples for Organization CreateApiKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Organization_CreateAPIKey_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organization_CreateAPIKey_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationCreateAPIKeyMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().createApiKeyWithResponse("rgconfluent", "pzvuoywx", "jqoxoahobqmhnklw", "ypyzlfhbml",
            new CreateApiKeyModel().withName("izlvofweryqgdgq").withDescription("vdxsmrddjlsfcsnwjezjraxgbkn"),
            com.azure.core.util.Context.NONE);
    }
}
