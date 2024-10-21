
import com.azure.resourcemanager.iotoperations.models.DataflowBuiltInTransformationSettings;
import com.azure.resourcemanager.iotoperations.models.DataflowDestinationOperationSettings;
import com.azure.resourcemanager.iotoperations.models.DataflowOperation;
import com.azure.resourcemanager.iotoperations.models.DataflowProperties;
import com.azure.resourcemanager.iotoperations.models.DataflowSourceOperationSettings;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import com.azure.resourcemanager.iotoperations.models.OperationType;
import com.azure.resourcemanager.iotoperations.models.OperationalMode;
import com.azure.resourcemanager.iotoperations.models.TransformationSerializationFormat;
import java.util.Arrays;

/**
 * Samples for Dataflow CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/Dataflow_CreateOrUpdate_SimpleFabric.json
     */
    /**
     * Sample code: Dataflow_CreateOrUpdate_SimpleFabric.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowCreateOrUpdateSimpleFabric(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflows().define("aio-to-fabric")
            .withExistingDataflowProfile("rgiotoperations", "resource-name123", "resource-name123")
            .withExtendedLocation(
                new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv").withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(
                new DataflowProperties()
                    .withMode(OperationalMode.ENABLED).withOperations(Arrays.asList(
                        new DataflowOperation().withOperationType(OperationType.SOURCE).withName("source1")
                            .withSourceSettings(new DataflowSourceOperationSettings()
                                .withEndpointRef("aio-builtin-broker-endpoint")
                                .withDataSources(Arrays.asList("azure-iot-operations/data/thermostat"))),
                        new DataflowOperation().withOperationType(OperationType.BUILT_IN_TRANSFORMATION)
                            .withBuiltInTransformationSettings(new DataflowBuiltInTransformationSettings()
                                .withSerializationFormat(TransformationSerializationFormat.PARQUET)
                                .withSchemaRef("aio-sr://exampleNamespace/exmapleParquetSchema:1.0.0")),
                        new DataflowOperation().withOperationType(OperationType.DESTINATION).withName("destination1")
                            .withDestinationSettings(new DataflowDestinationOperationSettings()
                                .withEndpointRef("fabric-endpoint").withDataDestination("telemetryTable")))))
            .create();
    }
}
