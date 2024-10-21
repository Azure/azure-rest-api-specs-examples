
import com.azure.resourcemanager.iotoperations.models.DataflowBuiltInTransformationFilter;
import com.azure.resourcemanager.iotoperations.models.DataflowBuiltInTransformationMap;
import com.azure.resourcemanager.iotoperations.models.DataflowBuiltInTransformationSettings;
import com.azure.resourcemanager.iotoperations.models.DataflowDestinationOperationSettings;
import com.azure.resourcemanager.iotoperations.models.DataflowOperation;
import com.azure.resourcemanager.iotoperations.models.DataflowProperties;
import com.azure.resourcemanager.iotoperations.models.DataflowSourceOperationSettings;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import com.azure.resourcemanager.iotoperations.models.OperationType;
import com.azure.resourcemanager.iotoperations.models.OperationalMode;
import java.util.Arrays;

/**
 * Samples for Dataflow CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/Dataflow_CreateOrUpdate_ComplexEventHub.json
     */
    /**
     * Sample code: Dataflow_CreateOrUpdate_ComplexEventHub.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowCreateOrUpdateComplexEventHub(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflows().define("aio-to-event-hub-transformed")
            .withExistingDataflowProfile("rgiotoperations", "resource-name123", "resource-name123")
            .withExtendedLocation(new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(
                new DataflowProperties().withMode(OperationalMode.ENABLED)
                    .withOperations(
                        Arrays
                            .asList(
                                new DataflowOperation().withOperationType(OperationType.SOURCE).withName("source1")
                                    .withSourceSettings(new DataflowSourceOperationSettings()
                                        .withEndpointRef("aio-builtin-broker-endpoint").withDataSources(Arrays
                                            .asList("azure-iot-operations/data/thermostat"))),
                                new DataflowOperation().withOperationType(OperationType.BUILT_IN_TRANSFORMATION)
                                    .withBuiltInTransformationSettings(new DataflowBuiltInTransformationSettings()
                                        .withFilter(Arrays.asList(new DataflowBuiltInTransformationFilter()
                                            .withInputs(Arrays.asList("temperature.Value", "\"Tag 10\".Value"))
                                            .withExpression("$1 > 9000 && $2 >= 8000")))
                                        .withMap(Arrays.asList(
                                            new DataflowBuiltInTransformationMap().withInputs(Arrays.asList("*"))
                                                .withOutput("*"),
                                            new DataflowBuiltInTransformationMap()
                                                .withInputs(Arrays.asList("temperature.Value", "\"Tag 10\".Value"))
                                                .withExpression("($1+$2)/2").withOutput("AvgTemp.Value"),
                                            new DataflowBuiltInTransformationMap().withInputs(Arrays.asList())
                                                .withExpression("true").withOutput("dataflow-processed"),
                                            new DataflowBuiltInTransformationMap()
                                                .withInputs(Arrays.asList("temperature.SourceTimestamp"))
                                                .withExpression("").withOutput(""),
                                            new DataflowBuiltInTransformationMap()
                                                .withInputs(Arrays.asList("\"Tag 10\"")).withExpression("")
                                                .withOutput("pressure"),
                                            new DataflowBuiltInTransformationMap()
                                                .withInputs(Arrays.asList("temperature.Value"))
                                                .withExpression("cToF($1)").withOutput("temperatureF.Value"),
                                            new DataflowBuiltInTransformationMap()
                                                .withInputs(Arrays.asList("\"Tag 10\".Value"))
                                                .withExpression("scale ($1,0,10,0,100)")
                                                .withOutput("\"Scale Tag 10\".Value")))),
                                new DataflowOperation().withOperationType(OperationType.DESTINATION)
                                    .withName("destination1").withDestinationSettings(
                                        new DataflowDestinationOperationSettings().withEndpointRef("event-hub-endpoint")
                                            .withDataDestination("myuniqueeventhub")))))
            .create();
    }
}
