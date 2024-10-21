
import com.azure.resourcemanager.iotoperations.models.DataflowBuiltInTransformationDataset;
import com.azure.resourcemanager.iotoperations.models.DataflowBuiltInTransformationFilter;
import com.azure.resourcemanager.iotoperations.models.DataflowBuiltInTransformationMap;
import com.azure.resourcemanager.iotoperations.models.DataflowBuiltInTransformationSettings;
import com.azure.resourcemanager.iotoperations.models.DataflowDestinationOperationSettings;
import com.azure.resourcemanager.iotoperations.models.DataflowMappingType;
import com.azure.resourcemanager.iotoperations.models.DataflowOperation;
import com.azure.resourcemanager.iotoperations.models.DataflowProperties;
import com.azure.resourcemanager.iotoperations.models.DataflowSourceOperationSettings;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import com.azure.resourcemanager.iotoperations.models.FilterType;
import com.azure.resourcemanager.iotoperations.models.OperationType;
import com.azure.resourcemanager.iotoperations.models.OperationalMode;
import com.azure.resourcemanager.iotoperations.models.SourceSerializationFormat;
import com.azure.resourcemanager.iotoperations.models.TransformationSerializationFormat;
import java.util.Arrays;

/**
 * Samples for Dataflow CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/Dataflow_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: Dataflow_CreateOrUpdate.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void dataflowCreateOrUpdate(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflows().define("resource-name123")
            .withExistingDataflowProfile("rgiotoperations", "resource-name123", "resource-name123")
            .withExtendedLocation(
                new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv").withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(new DataflowProperties().withMode(OperationalMode.ENABLED).withOperations(
                Arrays.asList(new DataflowOperation().withOperationType(OperationType.SOURCE).withName("knnafvkwoeakm")
                    .withSourceSettings(new DataflowSourceOperationSettings()
                        .withEndpointRef("iixotodhvhkkfcfyrkoveslqig").withAssetRef("zayyykwmckaocywdkohmu")
                        .withSerializationFormat(SourceSerializationFormat.JSON).withSchemaRef("pknmdzqll")
                        .withDataSources(Arrays.asList("chkkpymxhp")))
                    .withBuiltInTransformationSettings(new DataflowBuiltInTransformationSettings()
                        .withSerializationFormat(TransformationSerializationFormat.DELTA).withSchemaRef("mcdc")
                        .withDatasets(Arrays.asList(new DataflowBuiltInTransformationDataset()
                            .withKey("fakeTokenPlaceholder")
                            .withDescription("Lorem ipsum odor amet, consectetuer adipiscing elit.").withSchemaRef("n")
                            .withInputs(Arrays.asList("mosffpsslifkq")).withExpression("aatbwomvflemsxialv")))
                        .withFilter(Arrays.asList(new DataflowBuiltInTransformationFilter().withType(FilterType.FILTER)
                            .withDescription("Lorem ipsum odor amet, consectetuer adipiscing elit.")
                            .withInputs(Arrays.asList("sxmjkbntgb")).withExpression("n")))
                        .withMap(Arrays
                            .asList(new DataflowBuiltInTransformationMap().withType(DataflowMappingType.NEW_PROPERTIES)
                                .withDescription("Lorem ipsum odor amet, consectetuer adipiscing elit.")
                                .withInputs(Arrays.asList("xsbxuk")).withExpression("txoiltogsarwkzalsphvlmt")
                                .withOutput("nvgtmkfl"))))
                    .withDestinationSettings(new DataflowDestinationOperationSettings()
                        .withEndpointRef("kybkchnzimerguekuvqlqiqdvvrt").withDataDestination("cbrh")))))
            .create();
    }
}
