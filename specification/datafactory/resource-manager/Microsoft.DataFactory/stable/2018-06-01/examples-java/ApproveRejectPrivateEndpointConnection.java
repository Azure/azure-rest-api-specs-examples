
import com.azure.resourcemanager.datafactory.models.PrivateEndpoint;
import com.azure.resourcemanager.datafactory.models.PrivateLinkConnectionApprovalRequest;
import com.azure.resourcemanager.datafactory.models.PrivateLinkConnectionState;

/**
 * Samples for PrivateEndpointConnectionOperation CreateOrUpdateSync.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * ApproveRejectPrivateEndpointConnection.json
     */
    /**
     * Sample code: Approves or rejects a private endpoint connection for a factory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void approvesOrRejectsAPrivateEndpointConnectionForAFactory(
        com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.privateEndpointConnectionOperations().define("connection")
            .withExistingFactory("exampleResourceGroup", "exampleFactoryName")
            .withProperties(new PrivateLinkConnectionApprovalRequest()
                .withPrivateLinkServiceConnectionState(new PrivateLinkConnectionState().withStatus("Approved")
                    .withDescription("Approved by admin.").withActionsRequired(""))
                .withPrivateEndpoint(new PrivateEndpoint().withId(
                    "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/privateEndpoints/myPrivateEndpoint")))
            .create();
    }
}
