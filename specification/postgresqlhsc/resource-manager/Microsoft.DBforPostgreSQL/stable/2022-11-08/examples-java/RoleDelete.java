/** Samples for Roles Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/RoleDelete.json
     */
    /**
     * Sample code: RoleDelete.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void roleDelete(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.roles().delete("TestGroup", "pgtestsvc4", "role1", com.azure.core.util.Context.NONE);
    }
}
