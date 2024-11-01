
import com.azure.resourcemanager.streamanalytics.models.AuthenticationMode;
import com.azure.resourcemanager.streamanalytics.models.PostgreSqlOutputDataSource;

/**
 * Samples for Outputs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Create_PostgreSQL.json
     */
    /**
     * Sample code: Create a PostgreSQL output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        createAPostgreSQLOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().define("output3022").withExistingStreamingjob("sjrg7983", "sj2331")
            .withDatasource(new PostgreSqlOutputDataSource().withServer("someServer").withDatabase("someDatabase")
                .withTable("someTable").withUser("user").withPassword("fakeTokenPlaceholder").withMaxWriterCount(1.0F)
                .withAuthenticationMode(AuthenticationMode.MSI))
            .create();
    }
}
