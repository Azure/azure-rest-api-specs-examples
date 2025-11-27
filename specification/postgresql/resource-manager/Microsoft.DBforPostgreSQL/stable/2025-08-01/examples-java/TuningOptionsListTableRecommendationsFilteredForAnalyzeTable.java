
import com.azure.resourcemanager.postgresqlflexibleserver.models.RecommendationTypeParameterEnum;
import com.azure.resourcemanager.postgresqlflexibleserver.models.TuningOptionParameterEnum;

/**
 * Samples for TuningOptionsOperation ListRecommendations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * TuningOptionsListTableRecommendationsFilteredForAnalyzeTable.json
     */
    /**
     * Sample code: List available table recommendations, filtered to exclusively get those of ANALYZE TABLE type.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listAvailableTableRecommendationsFilteredToExclusivelyGetThoseOfANALYZETABLEType(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.tuningOptionsOperations().listRecommendations("exampleresourcegroup", "exampleserver",
            TuningOptionParameterEnum.TABLE, RecommendationTypeParameterEnum.ANALYZE_TABLE,
            com.azure.core.util.Context.NONE);
    }
}
