import com.azure.resourcemanager.kusto.models.PrivateLinkServiceConnectionStateProperty;

/** Samples for PrivateEndpointConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoPrivateEndpointConnectionsCreateOrUpdate.json
     */
    /**
     * Sample code: Approve or reject a private endpoint connection with a given name.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void approveOrRejectAPrivateEndpointConnectionWithAGivenName(
        com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .privateEndpointConnections()
            .define("privateEndpointTest")
            .withExistingCluster("kustorptest", "kustoclusterrptest4")
            .withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionStateProperty()
                    .withStatus("Approved")
                    .withDescription("Approved by johndoe@contoso.com"))
            .create();
    }
}
