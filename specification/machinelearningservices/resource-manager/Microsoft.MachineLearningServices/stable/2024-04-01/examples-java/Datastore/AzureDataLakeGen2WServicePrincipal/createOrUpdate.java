
import com.azure.resourcemanager.machinelearning.models.AzureDataLakeGen2Datastore;
import com.azure.resourcemanager.machinelearning.models.ServicePrincipalDatastoreCredentials;
import com.azure.resourcemanager.machinelearning.models.ServicePrincipalDatastoreSecrets;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

/**
 * Samples for Datastores CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Datastore/AzureDataLakeGen2WServicePrincipal/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate datastore (Azure Data Lake Gen2 w/ Service Principal).
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateDatastoreAzureDataLakeGen2WServicePrincipal(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.datastores().define("string").withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(
                new AzureDataLakeGen2Datastore().withDescription("string").withTags(mapOf("string", "string"))
                    .withCredentials(new ServicePrincipalDatastoreCredentials().withAuthorityUrl("string")
                        .withResourceUrl("string").withTenantId(UUID.fromString("00000000-1111-2222-3333-444444444444"))
                        .withClientId(UUID.fromString("00000000-1111-2222-3333-444444444444"))
                        .withSecrets(new ServicePrincipalDatastoreSecrets().withClientSecret("fakeTokenPlaceholder")))
                    .withFilesystem("string").withAccountName("string").withEndpoint("string").withProtocol("string"))
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
