```java
import java.util.HashMap;
import java.util.Map;

/** Samples for LoadTests CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/preview/2021-12-01-preview/examples/LoadTests_CreateOrUpdate.json
     */
    /**
     * Sample code: LoadTests_CreateOrUpdate.
     *
     * @param manager Entry point to LoadTestManager.
     */
    public static void loadTestsCreateOrUpdate(com.azure.resourcemanager.loadtestservice.LoadTestManager manager) {
        manager
            .loadTests()
            .define("myLoadTest")
            .withRegion("westus")
            .withExistingResourceGroup("dummyrg")
            .withTags(mapOf("Team", "Dev Exp"))
            .withDescription("This is new load test resource")
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-loadtestservice_1.0.0-beta.1/sdk/loadtestservice/azure-resourcemanager-loadtestservice/README.md) on how to add the SDK to your project and authenticate.
