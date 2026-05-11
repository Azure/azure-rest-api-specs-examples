
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
     * x-ms-original-file: 2026-03-01/Dataflow_CreateOrUpdate_SimpleEventGrid.json
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
            .withProperties(new DataflowProperties().withMode(OperationalMode.ENABLED).withOperations(Arrays.asList(
                new DataflowOperation().withOperationType(OperationType.SOURCE).withName("source1").withSourceSettings(
                    new DataflowSourceOperationSettings().withEndpointRef("aio-builtin-broker-endpoint")
                        .withDataSources(Arrays.asList("thermostats/+/telemetry/temperature/#"))),
                new DataflowOperation().withOperationType(OperationType.DESTINATION).withName("destination1")
                    .withDestinationSettings(new DataflowDestinationOperationSettings()
                        .withEndpointRef("event-grid-endpoint").withDataDestination("factory/telemetry")))))
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
            .create();
    }
}
