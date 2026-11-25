
import com.azure.resourcemanager.machinelearning.models.SecretExpiry;

/**
 * Samples for Datastores ListSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Datastore/listSecrets.json
     */
    /**
     * Sample code: Get datastore secrets.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getDatastoreSecrets(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.datastores().listSecretsWithResponse("test-rg", "my-aml-workspace", "string",
            new SecretExpiry().withExpirableSecret(false).withExpireAfterHours(1), com.azure.core.util.Context.NONE);
    }
}
