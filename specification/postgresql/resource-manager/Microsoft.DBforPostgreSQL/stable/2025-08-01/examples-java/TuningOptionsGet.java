
import com.azure.resourcemanager.postgresqlflexibleserver.models.TuningOptionParameterEnum;

/**
 * Samples for TuningOptionsOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/TuningOptionsGet.
     * json
     */
    /**
     * Sample code: Get the tuning options of a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        getTheTuningOptionsOfAServer(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.tuningOptionsOperations().getWithResponse("exampleresourcegroup", "exampleserver",
            TuningOptionParameterEnum.INDEX, com.azure.core.util.Context.NONE);
    }
}
