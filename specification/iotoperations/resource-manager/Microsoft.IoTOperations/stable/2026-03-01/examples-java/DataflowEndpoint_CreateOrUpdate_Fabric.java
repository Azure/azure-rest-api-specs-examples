
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointAuthenticationSystemAssignedManagedIdentity;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointFabricOneLake;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointFabricOneLakeAuthentication;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointFabricOneLakeNames;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointFabricPathType;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointProperties;
import com.azure.resourcemanager.iotoperations.models.EndpointType;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import com.azure.resourcemanager.iotoperations.models.FabricOneLakeAuthMethod;

/**
 * Samples for DataflowEndpoint CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/DataflowEndpoint_CreateOrUpdate_Fabric.json
     */
    /**
     * Sample code: DataflowEndpoint_CreateOrUpdate_Fabric.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowEndpointCreateOrUpdateFabric(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowEndpoints().define("fabric-endpoint")
            .withExistingInstance("rgiotoperations", "resource-name123")
            .withProperties(new DataflowEndpointProperties().withEndpointType(EndpointType.FABRIC_ONE_LAKE)
                .withFabricOneLakeSettings(new DataflowEndpointFabricOneLake()
                    .withAuthentication(new DataflowEndpointFabricOneLakeAuthentication()
                        .withMethod(FabricOneLakeAuthMethod.SYSTEM_ASSIGNED_MANAGED_IDENTITY)
                        .withSystemAssignedManagedIdentitySettings(
                            new DataflowEndpointAuthenticationSystemAssignedManagedIdentity()))
                    .withNames(new DataflowEndpointFabricOneLakeNames().withLakehouseName("example-lakehouse")
                        .withWorkspaceName("example-workspace"))
                    .withOneLakePathType(DataflowEndpointFabricPathType.TABLES)
                    .withHost("onelake.dfs.fabric.microsoft.com")))
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
            .create();
    }
}
