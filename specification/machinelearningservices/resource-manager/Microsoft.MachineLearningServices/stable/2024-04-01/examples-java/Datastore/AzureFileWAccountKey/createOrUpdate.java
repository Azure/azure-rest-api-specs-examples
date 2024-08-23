
import com.azure.resourcemanager.machinelearning.models.AccountKeyDatastoreCredentials;
import com.azure.resourcemanager.machinelearning.models.AccountKeyDatastoreSecrets;
import com.azure.resourcemanager.machinelearning.models.AzureFileDatastore;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Datastores CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Datastore/AzureFileWAccountKey/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate datastore (Azure File store w/ AccountKey).
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateDatastoreAzureFileStoreWAccountKey(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.datastores().define("string").withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(new AzureFileDatastore().withDescription("string").withTags(mapOf("string", "string"))
                .withCredentials(new AccountKeyDatastoreCredentials()
                    .withSecrets(new AccountKeyDatastoreSecrets().withKey("fakeTokenPlaceholder")))
                .withAccountName("string").withFileShareName("string").withEndpoint("string").withProtocol("string"))
            .withSkipValidation(false).create();
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
