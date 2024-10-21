
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
     * x-ms-original-file: 2024-09-15-preview/Dataflow_CreateOrUpdate_SimpleEventGrid.json
     */
    /**
     * Sample code: Dataflow_CreateOrUpdate_SimpleEventGrid.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowCreateOrUpdateSimpleEventGrid(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflows().define("aio-to-event-grid")
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
                                        .withEndpointRef("aio-builtin-broker-endpoint")
                                        .withDataSources(Arrays.asList("thermostats/+/telemetry/temperature/#"))),
                                new DataflowOperation().withOperationType(OperationType.DESTINATION)
                                    .withName("destination1")
                                    .withDestinationSettings(new DataflowDestinationOperationSettings()
                                        .withEndpointRef("event-grid-endpoint")
                                        .withDataDestination("factory/telemetry")))))
            .create();
    }
}
