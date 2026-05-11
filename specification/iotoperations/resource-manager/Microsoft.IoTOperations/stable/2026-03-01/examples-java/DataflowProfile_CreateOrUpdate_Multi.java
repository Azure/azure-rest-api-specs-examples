
import com.azure.resourcemanager.iotoperations.models.DataflowProfileProperties;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;

/**
 * Samples for DataflowProfile CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/DataflowProfile_CreateOrUpdate_Multi.json
     */
    /**
     * Sample code: DataflowProfile_CreateOrUpdate_Multi.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowProfileCreateOrUpdateMulti(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowProfiles().define("aio-dataflowprofile")
            .withExistingInstance("rgiotoperations", "resource-name123")
            .withProperties(new DataflowProfileProperties().withInstanceCount(3))
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
            .create();
    }
}
