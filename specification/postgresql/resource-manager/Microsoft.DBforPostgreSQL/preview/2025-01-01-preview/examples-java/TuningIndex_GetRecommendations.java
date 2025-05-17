
import com.azure.resourcemanager.postgresqlflexibleserver.models.TuningOptionEnum;

/**
 * Samples for TuningIndex ListRecommendations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * TuningIndex_GetRecommendations.json
     */
    /**
     * Sample code: TuningIndex_ListRecommendations.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        tuningIndexListRecommendations(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.tuningIndexes().listRecommendations("testrg", "pgtestsvc4", TuningOptionEnum.INDEX, null,
            com.azure.core.util.Context.NONE);
    }
}
