import com.azure.core.util.Context;
import com.azure.resourcemanager.mysqlflexibleserver.models.ConfigurationForBatchUpdate;
import com.azure.resourcemanager.mysqlflexibleserver.models.ConfigurationListForBatchUpdate;
import java.util.Arrays;

/** Samples for Configurations BatchUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ConfigurationsBatchUpdate.json
     */
    /**
     * Sample code: ConfigurationList.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void configurationList(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager
            .configurations()
            .batchUpdate(
                "testrg",
                "mysqltestserver",
                new ConfigurationListForBatchUpdate()
                    .withValue(
                        Arrays
                            .asList(
                                new ConfigurationForBatchUpdate().withName("event_scheduler").withValue("OFF"),
                                new ConfigurationForBatchUpdate().withName("div_precision_increment").withValue("8"))),
                Context.NONE);
    }
}
