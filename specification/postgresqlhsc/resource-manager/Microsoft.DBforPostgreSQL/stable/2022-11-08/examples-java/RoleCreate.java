/** Samples for Roles Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/RoleCreate.json
     */
    /**
     * Sample code: RoleCreate.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void roleCreate(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager
            .roles()
            .define("role1")
            .withExistingServerGroupsv2("TestGroup", "pgtestsvc4")
            .withPassword("password")
            .create();
    }
}
