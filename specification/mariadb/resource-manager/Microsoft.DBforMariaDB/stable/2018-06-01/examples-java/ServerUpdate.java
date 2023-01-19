import com.azure.resourcemanager.mariadb.models.Server;
import com.azure.resourcemanager.mariadb.models.SslEnforcementEnum;

/** Samples for Servers Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerUpdate.json
     */
    /**
     * Sample code: ServerUpdate.
     *
     * @param manager Entry point to MariaDBManager.
     */
    public static void serverUpdate(com.azure.resourcemanager.mariadb.MariaDBManager manager) {
        Server resource =
            manager
                .servers()
                .getByResourceGroupWithResponse("testrg", "mariadbtestsvc4", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withAdministratorLoginPassword("<administratorLoginPassword>")
            .withSslEnforcement(SslEnforcementEnum.DISABLED)
            .apply();
    }
}
