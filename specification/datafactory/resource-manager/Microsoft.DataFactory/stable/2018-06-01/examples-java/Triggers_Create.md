Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.12/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.datafactory.models.PipelineReference;
import com.azure.resourcemanager.datafactory.models.RecurrenceFrequency;
import com.azure.resourcemanager.datafactory.models.ScheduleTrigger;
import com.azure.resourcemanager.datafactory.models.ScheduleTriggerRecurrence;
import com.azure.resourcemanager.datafactory.models.TriggerPipelineReference;
import java.io.IOException;
import java.time.OffsetDateTime;
import java.util.Arrays;
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
                new ScheduleTrigger()
                    .withPipelines(
                        Arrays
                            .asList(
                                new TriggerPipelineReference()
                                    .withPipelineReference(new PipelineReference().withReferenceName("examplePipeline"))
                                    .withParameters(
                                        mapOf(
                                            "OutputBlobNameList",
                                            SerializerFactory
                                                .createDefaultManagementSerializerAdapter()
                                                .deserialize(
                                                    "[\"exampleoutput.csv\"]",
                                                    Object.class,
                                                    SerializerEncoding.JSON)))))
                    .withRecurrence(
                        new ScheduleTriggerRecurrence()
                            .withFrequency(RecurrenceFrequency.MINUTE)
                            .withInterval(4)
                            .withStartTime(OffsetDateTime.parse("2018-06-16T00:39:13.8441801Z"))
                            .withEndTime(OffsetDateTime.parse("2018-06-16T00:55:13.8441801Z"))
                            .withTimeZone("UTC")
                            .withAdditionalProperties(mapOf())))
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
