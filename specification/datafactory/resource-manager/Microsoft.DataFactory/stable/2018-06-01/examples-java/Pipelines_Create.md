Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.8/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.datafactory.models.Activity;
import com.azure.resourcemanager.datafactory.models.ParameterSpecification;
import com.azure.resourcemanager.datafactory.models.ParameterType;
import com.azure.resourcemanager.datafactory.models.PipelineElapsedTimeMetricPolicy;
import com.azure.resourcemanager.datafactory.models.PipelinePolicy;
import com.azure.resourcemanager.datafactory.models.VariableSpecification;
import com.azure.resourcemanager.datafactory.models.VariableType;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Pipelines CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Create.json
     */
    /**
     * Sample code: Pipelines_Create.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void pipelinesCreate(com.azure.resourcemanager.datafactory.DataFactoryManager manager)
        throws IOException {
        manager
            .pipelines()
            .define("examplePipeline")
            .withExistingFactory("exampleResourceGroup", "exampleFactoryName")
            .withActivities(
                Arrays
                    .asList(
                        new Activity()
                            .withName("ExampleForeachActivity")
                            .withAdditionalProperties(
                                mapOf(
                                    "typeProperties",
                                    SerializerFactory
                                        .createDefaultManagementSerializerAdapter()
                                        .deserialize(
                                            "{\"activities\":[{\"name\":\"ExampleCopyActivity\",\"type\":\"Copy\",\"inputs\":[{\"type\":\"DatasetReference\",\"parameters\":{\"MyFileName\":\"examplecontainer.csv\",\"MyFolderPath\":\"examplecontainer\"},\"referenceName\":\"exampleDataset\"}],\"outputs\":[{\"type\":\"DatasetReference\",\"parameters\":{\"MyFileName\":{\"type\":\"Expression\",\"value\":\"@item()\"},\"MyFolderPath\":\"examplecontainer\"},\"referenceName\":\"exampleDataset\"}],\"typeProperties\":{\"dataIntegrationUnits\":32,\"sink\":{\"type\":\"BlobSink\"},\"source\":{\"type\":\"BlobSource\"}}}],\"isSequential\":true,\"items\":{\"type\":\"Expression\",\"value\":\"@pipeline().parameters.OutputBlobNameList\"}}",
                                            Object.class,
                                            SerializerEncoding.JSON),
                                    "type",
                                    "ForEach"))))
            .withParameters(
                mapOf(
                    "JobId",
                    new ParameterSpecification().withType(ParameterType.STRING),
                    "OutputBlobNameList",
                    new ParameterSpecification().withType(ParameterType.ARRAY)))
            .withVariables(mapOf("TestVariableArray", new VariableSpecification().withType(VariableType.ARRAY)))
            .withRunDimensions(
                mapOf(
                    "JobId",
                    SerializerFactory
                        .createDefaultManagementSerializerAdapter()
                        .deserialize(
                            "{\"type\":\"Expression\",\"value\":\"@pipeline().parameters.JobId\"}",
                            Object.class,
                            SerializerEncoding.JSON)))
            .withPolicy(
                new PipelinePolicy()
                    .withElapsedTimeMetric(new PipelineElapsedTimeMetricPolicy().withDuration("0.00:10:00")))
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
