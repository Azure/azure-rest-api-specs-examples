
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointLocalStorage;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointProperties;
import com.azure.resourcemanager.iotoperations.models.EndpointType;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;

/**
 * Samples for DataflowEndpoint CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/DataflowEndpoint_CreateOrUpdate_LocalStorage.json
     */
    /**
     * Sample code: DataflowEndpoint_CreateOrUpdate_LocalStorage.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void dataflowEndpointCreateOrUpdateLocalStorage(
        com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowEndpoints().define("local-storage-endpoint")
            .withExistingInstance("rgiotoperations", "resource-name123")
            .withProperties(
                new DataflowEndpointProperties().withEndpointType(EndpointType.LOCAL_STORAGE).withLocalStorageSettings(
                    new DataflowEndpointLocalStorage().withPersistentVolumeClaimRef("example-pvc")))
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
            .create();
    }
}
