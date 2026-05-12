
import com.azure.resourcemanager.iotoperations.models.BatchingConfiguration;
import com.azure.resourcemanager.iotoperations.models.DataExplorerAuthMethod;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointAuthenticationSystemAssignedManagedIdentity;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointDataExplorer;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointDataExplorerAuthentication;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointProperties;
import com.azure.resourcemanager.iotoperations.models.EndpointType;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;

/**
 * Samples for DataflowEndpoint CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/DataflowEndpoint_CreateOrUpdate_ADX.json
     */
    /**
     * Sample code: DataflowEndpoint_CreateOrUpdate_ADX.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowEndpointCreateOrUpdateADX(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowEndpoints().define("adx-endpoint").withExistingInstance("rgiotoperations", "resource-name123")
            .withProperties(new DataflowEndpointProperties().withEndpointType(EndpointType.DATA_EXPLORER)
                .withDataExplorerSettings(new DataflowEndpointDataExplorer()
                    .withAuthentication(new DataflowEndpointDataExplorerAuthentication()
                        .withMethod(DataExplorerAuthMethod.SYSTEM_ASSIGNED_MANAGED_IDENTITY)
                        .withSystemAssignedManagedIdentitySettings(
                            new DataflowEndpointAuthenticationSystemAssignedManagedIdentity()))
                    .withDatabase("example-database").withHost("example.westeurope.kusto.windows.net")
                    .withBatching(new BatchingConfiguration().withLatencySeconds(9312).withMaxMessages(9028))))
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
            .create();
    }
}
