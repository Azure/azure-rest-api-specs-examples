
import com.azure.resourcemanager.postgresqlflexibleserver.models.PrincipalType;

/**
 * Samples for Administrators Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-12-01-preview/examples/
     * AdministratorAdd.json
     */
    /**
     * Sample code: Adds an Active DIrectory Administrator for the server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void addsAnActiveDIrectoryAdministratorForTheServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.administrators().define("oooooooo-oooo-oooo-oooo-oooooooooooo")
            .withExistingFlexibleServer("testrg", "testserver").withPrincipalType(PrincipalType.USER)
            .withPrincipalName("testuser1@microsoft.com").withTenantId("tttttttt-tttt-tttt-tttt-tttttttttttt").create();
    }
}
