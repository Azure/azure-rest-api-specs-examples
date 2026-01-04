
import com.azure.resourcemanager.appservice.fluent.models.DatabaseConnectionInner;

/**
 * Samples for StaticSites CreateOrUpdateBuildDatabaseConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * CreateOrUpdateStaticSiteBuildDatabaseConnection.json
     */
    /**
     * Sample code: Create or update a database connection for a static site build.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateADatabaseConnectionForAStaticSiteBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().createOrUpdateBuildDatabaseConnectionWithResponse(
            "rg", "testStaticSite0", "default", "default",
            new DatabaseConnectionInner().withResourceId(
                "/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/databaseRG/providers/Microsoft.DocumentDB/databaseAccounts/exampleDatabaseName")
                .withConnectionIdentity("SystemAssigned")
                .withConnectionString(
                    "AccountEndpoint=https://exampleDatabaseName.documents.azure.com:443/;Database=mydb;")
                .withRegion("West US 2"),
            com.azure.core.util.Context.NONE);
    }
}
