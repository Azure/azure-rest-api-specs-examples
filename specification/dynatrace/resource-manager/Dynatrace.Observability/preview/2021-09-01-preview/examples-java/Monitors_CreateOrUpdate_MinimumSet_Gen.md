```java
import java.util.HashMap;
import java.util.Map;

/** Samples for Monitors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/Monitors_CreateOrUpdate_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_CreateOrUpdate_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsCreateOrUpdateMinimumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager
            .monitors()
            .define("myMonitor")
            .withRegion("West US 2")
            .withExistingResourceGroup("myResourceGroup")
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-dynatrace_1.0.0-beta.1/sdk/dynatrace/azure-resourcemanager-dynatrace/README.md) on how to add the SDK to your project and authenticate.
