
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
     * x-ms-original-file: 2024-11-01/DataflowEndpoint_CreateOrUpdate_ADX.json
     */
    /**
     * Sample code: DataflowEndpoint_CreateOrUpdate_ADX.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowEndpointCreateOrUpdateADX(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowEndpoints().define("adx-endpoint").withExistingInstance("rgiotoperations", "resource-name123")
            .withExtendedLocation(
                new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv").withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(new DataflowEndpointProperties().withEndpointType(EndpointType.DATA_EXPLORER)
                .withDataExplorerSettings(new DataflowEndpointDataExplorer()
                    .withAuthentication(new DataflowEndpointDataExplorerAuthentication()
                        .withMethod(DataExplorerAuthMethod.SYSTEM_ASSIGNED_MANAGED_IDENTITY)
                        .withSystemAssignedManagedIdentitySettings(
                            new DataflowEndpointAuthenticationSystemAssignedManagedIdentity()))
                    .withDatabase("example-database").withHost("example.westeurope.kusto.windows.net")
                    .withBatching(new BatchingConfiguration().withLatencySeconds(9312).withMaxMessages(9028))))
            .create();
    }
}
