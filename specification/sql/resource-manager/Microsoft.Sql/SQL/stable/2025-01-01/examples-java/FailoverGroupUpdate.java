
import com.azure.resourcemanager.sql.models.FailoverGroupReadWriteEndpoint;
import com.azure.resourcemanager.sql.models.FailoverGroupUpdate;
import com.azure.resourcemanager.sql.models.ReadWriteEndpointFailoverPolicy;
import java.util.Arrays;

/**
 * Samples for FailoverGroups Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/FailoverGroupUpdate.json
     */
    /**
     * Sample code: Update failover group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateFailoverGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getFailoverGroups()
            .update("Default", "failover-group-primary-server", "failover-group-test-1", new FailoverGroupUpdate()
                .withReadWriteEndpoint(
                    new FailoverGroupReadWriteEndpoint().withFailoverPolicy(ReadWriteEndpointFailoverPolicy.AUTOMATIC)
                        .withFailoverWithDataLossGracePeriodMinutes(120))
                .withDatabases(Arrays.asList(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-1")),
                com.azure.core.util.Context.NONE);
    }
}
