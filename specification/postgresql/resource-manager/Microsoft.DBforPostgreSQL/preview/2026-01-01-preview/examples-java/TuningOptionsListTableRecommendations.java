
import com.azure.resourcemanager.postgresqlflexibleserver.models.TuningOptionParameterEnum;

/**
 * Samples for TuningOptionsOperation ListRecommendations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/TuningOptionsListTableRecommendations.json
     */
    /**
     * Sample code: List available table recommendations.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listAvailableTableRecommendations(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.tuningOptionsOperations().listRecommendations("exampleresourcegroup", "exampleserver",
            TuningOptionParameterEnum.TABLE, null, com.azure.core.util.Context.NONE);
    }
}
