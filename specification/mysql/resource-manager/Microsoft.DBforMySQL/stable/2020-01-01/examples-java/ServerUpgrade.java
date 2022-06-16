import com.azure.core.util.Context;
import com.azure.resourcemanager.mysql.models.ServerUpgradeParameters;

/** Samples for Servers Upgrade. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2020-01-01/examples/ServerUpgrade.json
     */
    /**
     * Sample code: ServerUpgrade.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void serverUpgrade(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .servers()
            .upgrade(
                "TestGroup", "testserver", new ServerUpgradeParameters().withTargetServerVersion("5.7"), Context.NONE);
    }
}
