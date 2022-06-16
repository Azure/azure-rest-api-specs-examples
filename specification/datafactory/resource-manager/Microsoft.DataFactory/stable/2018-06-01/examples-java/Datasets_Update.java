import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.datafactory.models.AzureBlobDataset;
import com.azure.resourcemanager.datafactory.models.DatasetResource;
import com.azure.resourcemanager.datafactory.models.LinkedServiceReference;
import com.azure.resourcemanager.datafactory.models.ParameterSpecification;
import com.azure.resourcemanager.datafactory.models.ParameterType;
import com.azure.resourcemanager.datafactory.models.TextFormat;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/** Samples for Datasets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Update.json
     */
    /**
     * Sample code: Datasets_Update.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void datasetsUpdate(com.azure.resourcemanager.datafactory.DataFactoryManager manager)
        throws IOException {
        DatasetResource resource =
            manager
                .datasets()
                .getWithResponse("exampleResourceGroup", "exampleFactoryName", "exampleDataset", null, Context.NONE)
                .getValue();
        resource
            .update()
            .withProperties(
                new AzureBlobDataset()
                    .withDescription("Example description")
                    .withLinkedServiceName(new LinkedServiceReference().withReferenceName("exampleLinkedService"))
                    .withParameters(
                        mapOf(
                            "MyFileName",
                            new ParameterSpecification().withType(ParameterType.STRING),
                            "MyFolderPath",
                            new ParameterSpecification().withType(ParameterType.STRING)))
                    .withFolderPath(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize(
                                "{\"type\":\"Expression\",\"value\":\"@dataset().MyFolderPath\"}",
                                Object.class,
                                SerializerEncoding.JSON))
                    .withFileName(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize(
                                "{\"type\":\"Expression\",\"value\":\"@dataset().MyFileName\"}",
                                Object.class,
                                SerializerEncoding.JSON))
                    .withFormat(new TextFormat()))
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
