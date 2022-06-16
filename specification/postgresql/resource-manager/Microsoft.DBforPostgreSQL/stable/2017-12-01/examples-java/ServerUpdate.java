import com.azure.core.util.Context;
import com.azure.resourcemanager.postgresql.models.MinimalTlsVersionEnum;
import com.azure.resourcemanager.postgresql.models.Server;
import com.azure.resourcemanager.postgresql.models.SslEnforcementEnum;

/** Samples for Servers Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerUpdate.json
     */
    /**
     * Sample code: ServerUpdate.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverUpdate(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        Server resource =
            manager.servers().getByResourceGroupWithResponse("testrg", "pgtestsvc4", Context.NONE).getValue();
        resource
            .update()
            .withAdministratorLoginPassword("<administratorLoginPassword>")
            .withSslEnforcement(SslEnforcementEnum.ENABLED)
            .withMinimalTlsVersion(MinimalTlsVersionEnum.TLS1_2)
            .apply();
    }
}
