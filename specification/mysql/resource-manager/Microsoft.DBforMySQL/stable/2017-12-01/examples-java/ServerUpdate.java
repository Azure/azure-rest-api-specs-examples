import com.azure.core.util.Context;
import com.azure.resourcemanager.mysql.models.Server;
import com.azure.resourcemanager.mysql.models.SslEnforcementEnum;

/** Samples for Servers Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerUpdate.json
     */
    /**
     * Sample code: ServerUpdate.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void serverUpdate(com.azure.resourcemanager.mysql.MySqlManager manager) {
        Server resource =
            manager.servers().getByResourceGroupWithResponse("testrg", "mysqltestsvc4", Context.NONE).getValue();
        resource
            .update()
            .withAdministratorLoginPassword("<administratorLoginPassword>")
            .withSslEnforcement(SslEnforcementEnum.DISABLED)
            .apply();
    }
}
