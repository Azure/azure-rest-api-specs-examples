```java
import com.azure.resourcemanager.machinelearning.models.AccountKeyDatastoreCredentials;
import com.azure.resourcemanager.machinelearning.models.AccountKeyDatastoreSecrets;
import com.azure.resourcemanager.machinelearning.models.AzureBlobDatastore;
import java.util.HashMap;
import java.util.Map;

/** Samples for Datastores CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Datastore/AzureBlobWAccountKey/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate datastore (AzureBlob w/ AccountKey).
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void createOrUpdateDatastoreAzureBlobWAccountKey(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .datastores()
            .define("string")
            .withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(
                new AzureBlobDatastore()
                    .withDescription("string")
                    .withTags(mapOf("string", "string"))
                    .withCredentials(
                        new AccountKeyDatastoreCredentials()
                            .withSecrets(new AccountKeyDatastoreSecrets().withKey("string")))
                    .withAccountName("string")
                    .withContainerName("string")
                    .withEndpoint("core.windows.net")
                    .withProtocol("https"))
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.1/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.
