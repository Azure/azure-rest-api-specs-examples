import com.azure.resourcemanager.databricks.models.PrivateEndpointConnectionProperties;
import com.azure.resourcemanager.databricks.models.PrivateLinkServiceConnectionState;
import com.azure.resourcemanager.databricks.models.PrivateLinkServiceConnectionStatus;

/** Samples for PrivateEndpointConnections Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/PrivateEndpointConnectionsUpdate.json
     */
    /**
     * Sample code: Update a private endpoint connection.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void updateAPrivateEndpointConnection(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager
            .privateEndpointConnections()
            .define("myWorkspace.23456789-1111-1111-1111-111111111111")
            .withExistingWorkspace("myResourceGroup", "myWorkspace")
            .withProperties(
                new PrivateEndpointConnectionProperties()
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState()
                            .withStatus(PrivateLinkServiceConnectionStatus.APPROVED)
                            .withDescription("Approved by databricksadmin@contoso.com")))
            .create();
    }
}
