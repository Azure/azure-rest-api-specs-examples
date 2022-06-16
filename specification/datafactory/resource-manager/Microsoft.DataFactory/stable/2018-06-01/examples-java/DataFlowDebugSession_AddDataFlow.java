import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.datafactory.models.AzureBlobStorageLinkedService;
import com.azure.resourcemanager.datafactory.models.AzureBlobStorageLocation;
import com.azure.resourcemanager.datafactory.models.DataFlowDebugPackage;
import com.azure.resourcemanager.datafactory.models.DataFlowDebugPackageDebugSettings;
import com.azure.resourcemanager.datafactory.models.DataFlowDebugResource;
import com.azure.resourcemanager.datafactory.models.DataFlowSource;
import com.azure.resourcemanager.datafactory.models.DataFlowSourceSetting;
import com.azure.resourcemanager.datafactory.models.DatasetDebugResource;
import com.azure.resourcemanager.datafactory.models.DatasetReference;
import com.azure.resourcemanager.datafactory.models.DelimitedTextDataset;
import com.azure.resourcemanager.datafactory.models.LinkedServiceDebugResource;
import com.azure.resourcemanager.datafactory.models.LinkedServiceReference;
import com.azure.resourcemanager.datafactory.models.MappingDataFlow;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for DataFlowDebugSession AddDataFlow. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_AddDataFlow.json
     */
    /**
     * Sample code: DataFlowDebugSession_AddDataFlow.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void dataFlowDebugSessionAddDataFlow(com.azure.resourcemanager.datafactory.DataFactoryManager manager)
        throws IOException {
        manager
            .dataFlowDebugSessions()
            .addDataFlowWithResponse(
                "exampleResourceGroup",
                "exampleFactoryName",
                new DataFlowDebugPackage()
                    .withSessionId("f06ed247-9d07-49b2-b05e-2cb4a2fc871e")
                    .withDataFlow(
                        new DataFlowDebugResource()
                            .withName("dataflow1")
                            .withProperties(
                                new MappingDataFlow()
                                    .withSources(
                                        Arrays
                                            .asList(
                                                new DataFlowSource()
                                                    .withName("source1")
                                                    .withDataset(
                                                        new DatasetReference().withReferenceName("DelimitedText2"))))
                                    .withSinks(Arrays.asList())
                                    .withTransformations(Arrays.asList())
                                    .withScript(
                                        "\n\n"
                                            + "source(output(\n"
                                            + "\t\tColumn_1 as string\n"
                                            + "\t),\n"
                                            + "\tallowSchemaDrift: true,\n"
                                            + "\tvalidateSchema: false) ~> source1")))
                    .withDatasets(
                        Arrays
                            .asList(
                                new DatasetDebugResource()
                                    .withName("dataset1")
                                    .withProperties(
                                        new DelimitedTextDataset()
                                            .withSchema(
                                                SerializerFactory
                                                    .createDefaultManagementSerializerAdapter()
                                                    .deserialize(
                                                        "[{\"type\":\"String\"}]",
                                                        Object.class,
                                                        SerializerEncoding.JSON))
                                            .withLinkedServiceName(
                                                new LinkedServiceReference().withReferenceName("linkedService5"))
                                            .withAnnotations(Arrays.asList())
                                            .withLocation(
                                                new AzureBlobStorageLocation()
                                                    .withFileName("Ansiencoding.csv")
                                                    .withContainer("dataflow-sample-data"))
                                            .withColumnDelimiter(",")
                                            .withQuoteChar("\"")
                                            .withEscapeChar("\\")
                                            .withFirstRowAsHeader(true))))
                    .withLinkedServices(
                        Arrays
                            .asList(
                                new LinkedServiceDebugResource()
                                    .withName("linkedService1")
                                    .withProperties(
                                        new AzureBlobStorageLinkedService()
                                            .withAnnotations(Arrays.asList())
                                            .withConnectionString(
                                                "DefaultEndpointsProtocol=https;AccountName=<storageName>;EndpointSuffix=core.windows.net;")
                                            .withEncryptedCredential("<credential>"))))
                    .withDebugSettings(
                        new DataFlowDebugPackageDebugSettings()
                            .withSourceSettings(
                                Arrays
                                    .asList(
                                        new DataFlowSourceSetting()
                                            .withSourceName("source1")
                                            .withRowLimit(1000)
                                            .withAdditionalProperties(mapOf()),
                                        new DataFlowSourceSetting()
                                            .withSourceName("source2")
                                            .withRowLimit(222)
                                            .withAdditionalProperties(mapOf())))
                            .withParameters(mapOf("sourcePath", "Toy"))
                            .withDatasetParameters(
                                SerializerFactory
                                    .createDefaultManagementSerializerAdapter()
                                    .deserialize(
                                        "{\"Movies\":{\"path\":\"abc\"},\"Output\":{\"time\":\"def\"}}",
                                        Object.class,
                                        SerializerEncoding.JSON)))
                    .withAdditionalProperties(mapOf()),
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
