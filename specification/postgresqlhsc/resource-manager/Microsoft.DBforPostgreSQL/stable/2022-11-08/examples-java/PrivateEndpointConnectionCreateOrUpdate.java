import com.azure.resourcemanager.cosmosdbforpostgresql.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.cosmosdbforpostgresql.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/PrivateEndpointConnectionCreateOrUpdate.json
     */
    /**
     * Sample code: Approves or Rejects a Private Endpoint Connection with a given name.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void approvesOrRejectsAPrivateEndpointConnectionWithAGivenName(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager
            .privateEndpointConnections()
            .define("private-endpoint-connection-name")
            .withExistingServerGroupsv2("TestGroup", "testcluster")
            .withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState()
                    .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("Approved by johndoe@contoso.com"))
            .create();
    }
}
