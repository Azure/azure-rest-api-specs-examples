
import com.azure.resourcemanager.postgresqlflexibleserver.models.TuningOptionParameterEnum;

/**
 * Samples for TuningOptionsOperation ListRecommendations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/TuningOptionsListIndexRecommendations.json
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
