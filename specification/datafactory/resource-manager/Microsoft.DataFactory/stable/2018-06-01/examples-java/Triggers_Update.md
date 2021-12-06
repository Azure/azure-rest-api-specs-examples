Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.8/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.datafactory.models.Trigger;
import com.azure.resourcemanager.datafactory.models.TriggerResource;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/** Samples for Triggers CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Update.json
     */
    /**
     * Sample code: Triggers_Update.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void triggersUpdate(com.azure.resourcemanager.datafactory.DataFactoryManager manager)
        throws IOException {
        TriggerResource resource =
            manager
                .triggers()
                .getWithResponse("exampleResourceGroup", "exampleFactoryName", "exampleTrigger", null, Context.NONE)
                .getValue();
        resource
            .update()
            .withProperties(
                new Trigger()
                    .withDescription("Example description")
                    .withAdditionalProperties(
                        mapOf(
                            "typeProperties",
                            SerializerFactory
                                .createDefaultManagementSerializerAdapter()
                                .deserialize(
                                    "{\"recurrence\":{\"endTime\":\"2018-06-16T00:55:14.905167Z\",\"frequency\":\"Minute\",\"interval\":4,\"startTime\":\"2018-06-16T00:39:14.905167Z\",\"timeZone\":\"UTC\"}}",
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
            .apply();
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
