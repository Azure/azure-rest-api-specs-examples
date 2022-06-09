```java
import com.azure.resourcemanager.machinelearning.models.ComputeInstance;
import com.azure.resourcemanager.machinelearning.models.ComputeInstanceProperties;
import java.util.HashMap;
import java.util.Map;

/** Samples for Compute CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Compute/createOrUpdate/ComputeInstanceMinimal.json
     */
    /**
     * Sample code: Create an ComputeInstance Compute with minimal inputs.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void createAnComputeInstanceComputeWithMinimalInputs(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .computes()
            .define("compute123")
            .withExistingWorkspace("testrg123", "workspaces123")
            .withRegion("eastus")
            .withProperties(
                new ComputeInstance().withProperties(new ComputeInstanceProperties().withVmSize("STANDARD_NC6")))
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
