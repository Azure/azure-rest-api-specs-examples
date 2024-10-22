
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
     * x-ms-original-file: 2024-09-15-preview/DataflowEndpoint_CreateOrUpdate_LocalStorage.json
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
            .withExtendedLocation(
                new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv").withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(
                new DataflowEndpointProperties().withEndpointType(EndpointType.LOCAL_STORAGE).withLocalStorageSettings(
                    new DataflowEndpointLocalStorage().withPersistentVolumeClaimRef("example-pvc")))
            .create();
    }
}
