Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.10/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/** Samples for Pipelines CreateRun. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_CreateRun.json
     */
    /**
     * Sample code: Pipelines_CreateRun.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void pipelinesCreateRun(com.azure.resourcemanager.datafactory.DataFactoryManager manager)
        throws IOException {
        manager
            .pipelines()
            .createRunWithResponse(
                "exampleResourceGroup",
                "exampleFactoryName",
                "examplePipeline",
                null,
                null,
                null,
                null,
                mapOf(
                    "OutputBlobNameList",
                    SerializerFactory
                        .createDefaultManagementSerializerAdapter()
                        .deserialize("[\"exampleoutput.csv\"]", Object.class, SerializerEncoding.JSON)),
                Context.NONE);
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
