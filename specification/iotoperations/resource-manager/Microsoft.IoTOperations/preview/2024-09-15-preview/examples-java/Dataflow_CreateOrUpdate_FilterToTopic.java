
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
import java.util.Arrays;

/**
 * Samples for Dataflow CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/Dataflow_CreateOrUpdate_FilterToTopic.json
     */
    /**
     * Sample code: Dataflow_CreateOrUpdate_FilterToTopic.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowCreateOrUpdateFilterToTopic(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflows().define("mqtt-filter-to-topic")
            .withExistingDataflowProfile("rgiotoperations", "resource-name123", "resource-name123")
            .withExtendedLocation(new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(
                new DataflowProperties().withMode(OperationalMode.ENABLED)
                    .withOperations(
                        Arrays.asList(
                            new DataflowOperation().withOperationType(OperationType.SOURCE).withName("source1")
                                .withSourceSettings(new DataflowSourceOperationSettings()
                                    .withEndpointRef("aio-builtin-broker-endpoint").withDataSources(
                                        Arrays.asList("azure-iot-operations/data/thermostat"))),
                            new DataflowOperation().withOperationType(OperationType.BUILT_IN_TRANSFORMATION)
                                .withName("transformation1")
                                .withBuiltInTransformationSettings(new DataflowBuiltInTransformationSettings()
                                    .withFilter(Arrays.asList(new DataflowBuiltInTransformationFilter()
                                        .withType(FilterType.FILTER).withDescription("filter-datapoint")
                                        .withInputs(Arrays.asList("temperature.Value", "\"Tag 10\".Value"))
                                        .withExpression("$1 > 9000 && $2 >= 8000")))
                                    .withMap(Arrays.asList(new DataflowBuiltInTransformationMap()
                                        .withType(DataflowMappingType.PASS_THROUGH).withInputs(Arrays.asList("*"))
                                        .withOutput("*")))),
                            new DataflowOperation().withOperationType(OperationType.DESTINATION)
                                .withName("destination1")
                                .withDestinationSettings(new DataflowDestinationOperationSettings()
                                    .withEndpointRef("aio-builtin-broker-endpoint")
                                    .withDataDestination("data/filtered/thermostat")))))
            .create();
    }
}
