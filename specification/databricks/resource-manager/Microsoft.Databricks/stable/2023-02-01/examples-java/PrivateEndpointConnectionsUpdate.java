import com.azure.resourcemanager.databricks.models.PrivateEndpointConnectionProperties;
import com.azure.resourcemanager.databricks.models.PrivateLinkServiceConnectionState;
import com.azure.resourcemanager.databricks.models.PrivateLinkServiceConnectionStatus;

/** Samples for PrivateEndpointConnections Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/PrivateEndpointConnectionsUpdate.json
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
