
/**
 * Samples for Roles ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-03-02-preview/RoleListByCluster.json
     */
    /**
     * Sample code: RoleList.
     * 
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void roleList(com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.roles().listByCluster("TestGroup", "pgtestsvc4", com.azure.core.util.Context.NONE);
    }
}
