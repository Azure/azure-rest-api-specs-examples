import com.azure.resourcemanager.machinelearning.models.AzureDataLakeGen1Datastore;
import com.azure.resourcemanager.machinelearning.models.ServicePrincipalDatastoreCredentials;
import com.azure.resourcemanager.machinelearning.models.ServicePrincipalDatastoreSecrets;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

/** Samples for Datastores CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Datastore/AzureDataLakeGen1WServicePrincipal/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate datastore (Azure Data Lake Gen1 w/ ServicePrincipal).
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void createOrUpdateDatastoreAzureDataLakeGen1WServicePrincipal(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .datastores()
            .define("string")
            .withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(
                new AzureDataLakeGen1Datastore()
                    .withDescription("string")
                    .withTags(mapOf("string", "string"))
                    .withCredentials(
                        new ServicePrincipalDatastoreCredentials()
                            .withAuthorityUrl("string")
                            .withClientId(UUID.fromString("00000000-1111-2222-3333-444444444444"))
                            .withResourceUrl("string")
                            .withSecrets(new ServicePrincipalDatastoreSecrets().withClientSecret("string"))
                            .withTenantId(UUID.fromString("00000000-1111-2222-3333-444444444444")))
                    .withStoreName("string"))
            .withSkipValidation(false)
            .create();
    }

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
