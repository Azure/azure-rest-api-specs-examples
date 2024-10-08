
import com.azure.resourcemanager.sql.fluent.models.RecommendedActionInner;
import com.azure.resourcemanager.sql.models.RecommendedActionCurrentState;
import com.azure.resourcemanager.sql.models.RecommendedActionStateInfo;

/**
 * Samples for DatabaseRecommendedActions Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseRecommendedActionUpdate.json
     */
    /**
     * Sample code: Update database recommended action.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateDatabaseRecommendedAction(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseRecommendedActions().updateWithResponse(
            "workloadinsight-demos", "misosisvr", "IndexAdvisor_test_3", "CreateIndex",
            "IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB",
            new RecommendedActionInner()
                .withState(new RecommendedActionStateInfo().withCurrentValue(RecommendedActionCurrentState.PENDING)),
            com.azure.core.util.Context.NONE);
    }
}
