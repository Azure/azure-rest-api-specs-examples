
import com.azure.resourcemanager.postgresqlflexibleserver.models.TuningOptionParameterEnum;

/**
 * Samples for TuningOptionsOperation ListRecommendations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * TuningOptionsListIndexRecommendations.json
     */
    /**
     * Sample code: List available index recommendations.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listAvailableIndexRecommendations(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.tuningOptionsOperations().listRecommendations("exampleresourcegroup", "exampleserver",
            TuningOptionParameterEnum.INDEX, null, com.azure.core.util.Context.NONE);
    }
}
