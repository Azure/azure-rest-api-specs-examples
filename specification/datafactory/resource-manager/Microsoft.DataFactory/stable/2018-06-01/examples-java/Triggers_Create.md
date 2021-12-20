Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.9/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.datafactory.models.Trigger;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/** Samples for Triggers CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Create.json
     */
    /**
     * Sample code: Triggers_Create.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void triggersCreate(com.azure.resourcemanager.datafactory.DataFactoryManager manager)
        throws IOException {
        manager
            .triggers()
            .define("exampleTrigger")
            .withExistingFactory("exampleResourceGroup", "exampleFactoryName")
            .withProperties(
                new Trigger()
                    .withAdditionalProperties(
                        mapOf(
                            "typeProperties",
                            SerializerFactory
                                .createDefaultManagementSerializerAdapter()
                                .deserialize(
                                    "{\"recurrence\":{\"endTime\":\"2018-06-16T00:55:13.8441801Z\",\"frequency\":\"Minute\",\"interval\":4,\"startTime\":\"2018-06-16T00:39:13.8441801Z\",\"timeZone\":\"UTC\"}}",
                                    Object.class,
                                    SerializerEncoding.JSON),
                            "pipelines",
                            SerializerFactory
                                .createDefaultManagementSerializerAdapter()
                                .deserialize(
                                    "[{\"parameters\":{\"OutputBlobNameList\":[\"exampleoutput.csv\"]},\"pipelineReference\":{\"type\":\"PipelineReference\",\"referenceName\":\"examplePipeline\"}}]",
                                    Object.class,
                                    SerializerEncoding.JSON),
                            "type",
                            "ScheduleTrigger")))
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
