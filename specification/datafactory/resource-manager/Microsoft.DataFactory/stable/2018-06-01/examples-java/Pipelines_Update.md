Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.13/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.datafactory.models.BlobSink;
import com.azure.resourcemanager.datafactory.models.BlobSource;
import com.azure.resourcemanager.datafactory.models.CopyActivity;
import com.azure.resourcemanager.datafactory.models.DatasetReference;
import com.azure.resourcemanager.datafactory.models.Expression;
import com.azure.resourcemanager.datafactory.models.ForEachActivity;
import com.azure.resourcemanager.datafactory.models.ParameterSpecification;
import com.azure.resourcemanager.datafactory.models.ParameterType;
import com.azure.resourcemanager.datafactory.models.PipelineElapsedTimeMetricPolicy;
import com.azure.resourcemanager.datafactory.models.PipelinePolicy;
import com.azure.resourcemanager.datafactory.models.PipelineResource;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Pipelines CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Update.json
     */
    /**
     * Sample code: Pipelines_Update.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void pipelinesUpdate(com.azure.resourcemanager.datafactory.DataFactoryManager manager)
        throws IOException {
        PipelineResource resource =
            manager
                .pipelines()
                .getWithResponse("exampleResourceGroup", "exampleFactoryName", "examplePipeline", null, Context.NONE)
                .getValue();
        resource
            .update()
            .withDescription("Example description")
            .withActivities(
                Arrays
                    .asList(
                        new ForEachActivity()
                            .withName("ExampleForeachActivity")
                            .withIsSequential(true)
                            .withItems(new Expression().withValue("@pipeline().parameters.OutputBlobNameList"))
                            .withActivities(
                                Arrays
                                    .asList(
                                        new CopyActivity()
                                            .withName("ExampleCopyActivity")
                                            .withInputs(
                                                Arrays
                                                    .asList(
                                                        new DatasetReference()
                                                            .withReferenceName("exampleDataset")
                                                            .withParameters(
                                                                mapOf(
                                                                    "MyFileName",
                                                                    "examplecontainer.csv",
                                                                    "MyFolderPath",
                                                                    "examplecontainer"))))
                                            .withOutputs(
                                                Arrays
                                                    .asList(
                                                        new DatasetReference()
                                                            .withReferenceName("exampleDataset")
                                                            .withParameters(
                                                                mapOf(
                                                                    "MyFileName",
                                                                    SerializerFactory
                                                                        .createDefaultManagementSerializerAdapter()
                                                                        .deserialize(
                                                                            "{\"type\":\"Expression\",\"value\":\"@item()\"}",
                                                                            Object.class,
                                                                            SerializerEncoding.JSON),
                                                                    "MyFolderPath",
                                                                    "examplecontainer"))))
                                            .withSource(new BlobSource())
                                            .withSink(new BlobSink())
                                            .withDataIntegrationUnits(32)))))
            .withParameters(mapOf("OutputBlobNameList", new ParameterSpecification().withType(ParameterType.ARRAY)))
            .withPolicy(
                new PipelinePolicy()
                    .withElapsedTimeMetric(new PipelineElapsedTimeMetricPolicy().withDuration("0.00:10:00")))
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
