/** Samples for Roles Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/RoleGet.json
     */
    /**
     * Sample code: Get the role of the cluster.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void getTheRoleOfTheCluster(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.roles().getWithResponse("TestGroup", "pgtestsvc4", "role1", com.azure.core.util.Context.NONE);
    }
}
