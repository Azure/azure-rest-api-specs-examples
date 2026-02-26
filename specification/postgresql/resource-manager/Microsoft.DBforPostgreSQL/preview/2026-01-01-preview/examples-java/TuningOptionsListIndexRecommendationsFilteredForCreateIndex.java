
import com.azure.resourcemanager.postgresqlflexibleserver.models.RecommendationTypeParameterEnum;
import com.azure.resourcemanager.postgresqlflexibleserver.models.TuningOptionParameterEnum;

/**
 * Samples for TuningOptionsOperation ListRecommendations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/TuningOptionsListIndexRecommendationsFilteredForCreateIndex.json
     */
    /**
     * Sample code: List available index recommendations, filtered to exclusively get those of CREATE INDEX type.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listAvailableIndexRecommendationsFilteredToExclusivelyGetThoseOfCREATEINDEXType(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.tuningOptionsOperations().listRecommendations("exampleresourcegroup", "exampleserver",
            TuningOptionParameterEnum.INDEX, RecommendationTypeParameterEnum.CREATE_INDEX,
            com.azure.core.util.Context.NONE);
    }
}
