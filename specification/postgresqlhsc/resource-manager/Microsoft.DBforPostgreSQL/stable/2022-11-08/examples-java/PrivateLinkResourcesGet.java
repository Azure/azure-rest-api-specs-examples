/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/PrivateLinkResourcesGet.json
     */
    /**
     * Sample code: Gets a private link resource for cluster.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void getsAPrivateLinkResourceForCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager
            .privateLinkResources()
            .getWithResponse("TestGroup", "testcluster", "plr", com.azure.core.util.Context.NONE);
    }
}
