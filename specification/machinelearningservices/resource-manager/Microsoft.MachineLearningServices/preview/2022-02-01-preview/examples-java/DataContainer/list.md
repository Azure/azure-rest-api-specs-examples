```java
import com.azure.core.util.Context;

/** Samples for DataContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/DataContainer/list.json
     */
    /**
     * Sample code: List Data Container.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listDataContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.dataContainers().list("testrg123", "workspace123", null, null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.1/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.
