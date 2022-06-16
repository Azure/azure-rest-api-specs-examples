import com.azure.core.util.Context;
import com.azure.resourcemanager.mysql.fluent.models.ConfigurationListResultInner;

/** Samples for ServerParameters ListUpdateConfigurations. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ConfigurationsUpdateByServer.json
     */
    /**
     * Sample code: ConfigurationList.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void configurationList(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .serverParameters()
            .listUpdateConfigurations("testrg", "mysqltestsvc1", new ConfigurationListResultInner(), Context.NONE);
    }
}
