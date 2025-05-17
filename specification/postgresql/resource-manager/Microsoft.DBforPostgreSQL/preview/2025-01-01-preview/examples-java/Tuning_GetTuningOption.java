
import com.azure.resourcemanager.postgresqlflexibleserver.models.TuningOptionEnum;

/**
 * Samples for TuningOptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * Tuning_GetTuningOption.json
     */
    /**
     * Sample code: TuningOptions_Get.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void tuningOptionsGet(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.tuningOptions().getWithResponse("testrg", "testserver", TuningOptionEnum.INDEX,
            com.azure.core.util.Context.NONE);
    }
}
