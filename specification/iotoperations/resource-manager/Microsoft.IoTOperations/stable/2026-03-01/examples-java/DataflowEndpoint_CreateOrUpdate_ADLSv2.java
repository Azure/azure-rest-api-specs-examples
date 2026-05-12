
import com.azure.resourcemanager.iotoperations.models.DataLakeStorageAuthMethod;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointAuthenticationAccessToken;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointDataLakeStorage;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointDataLakeStorageAuthentication;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointProperties;
import com.azure.resourcemanager.iotoperations.models.EndpointType;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;

/**
 * Samples for DataflowEndpoint CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/DataflowEndpoint_CreateOrUpdate_ADLSv2.json
     */
    /**
     * Sample code: DataflowEndpoint_CreateOrUpdate_ADLSv2.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowEndpointCreateOrUpdateADLSv2(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager
            .dataflowEndpoints().define(
                "adlsv2-endpoint")
            .withExistingInstance("rgiotoperations",
                "resource-name123")
            .withProperties(
                new DataflowEndpointProperties()
                    .withEndpointType(
                        EndpointType.DATA_LAKE_STORAGE)
                    .withDataLakeStorageSettings(new DataflowEndpointDataLakeStorage()
                        .withAuthentication(new DataflowEndpointDataLakeStorageAuthentication()
                            .withMethod(DataLakeStorageAuthMethod.ACCESS_TOKEN).withAccessTokenSettings(
                                new DataflowEndpointAuthenticationAccessToken().withSecretRef("fakeTokenPlaceholder")))
                        .withHost("example.blob.core.windows.net")))
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
            .create();
    }
}
