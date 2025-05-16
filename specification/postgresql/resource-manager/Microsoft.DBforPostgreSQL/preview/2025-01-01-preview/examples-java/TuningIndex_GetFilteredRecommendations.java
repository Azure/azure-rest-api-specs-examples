
import com.azure.resourcemanager.postgresqlflexibleserver.models.RecommendationType;
import com.azure.resourcemanager.postgresqlflexibleserver.models.TuningOptionEnum;

/**
 * Samples for TuningIndex ListRecommendations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * TuningIndex_GetFilteredRecommendations.json
     */
    /**
     * Sample code: TuningIndex_ListFilteredRecommendations.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void tuningIndexListFilteredRecommendations(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.tuningIndexes().listRecommendations("testrg", "pgtestrecs", TuningOptionEnum.INDEX,
            RecommendationType.CREATE_INDEX, com.azure.core.util.Context.NONE);
    }
}
