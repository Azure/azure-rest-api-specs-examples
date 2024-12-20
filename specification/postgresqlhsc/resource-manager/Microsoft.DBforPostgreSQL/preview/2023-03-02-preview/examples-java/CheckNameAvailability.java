
import com.azure.resourcemanager.cosmosdbforpostgresql.models.NameAvailabilityRequest;

/**
 * Samples for Clusters CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-02-preview/examples/
     * CheckNameAvailability.json
     */
    /**
     * Sample code: Check name availability.
     * 
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void
        checkNameAvailability(com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        manager.clusters().checkNameAvailabilityWithResponse(new NameAvailabilityRequest().withName("name1"),
            com.azure.core.util.Context.NONE);
    }
}
