```java
import com.azure.core.util.Context;

/** Samples for Datastores ListSecrets. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Datastore/listSecrets.json
     */
    /**
     * Sample code: Get datastore secrets.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getDatastoreSecrets(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.datastores().listSecretsWithResponse("test-rg", "my-aml-workspace", "string", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.1/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.
