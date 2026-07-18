
/**
 * Samples for DatabaseRecommendedActions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseRecommendedActionGet.json
     */
    /**
     * Sample code: Get database recommended action.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getDatabaseRecommendedAction(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseRecommendedActions().getWithResponse("workloadinsight-demos", "misosisvr",
            "IndexAdvisor_test_3", "CreateIndex", "IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB",
            com.azure.core.util.Context.NONE);
    }
}
