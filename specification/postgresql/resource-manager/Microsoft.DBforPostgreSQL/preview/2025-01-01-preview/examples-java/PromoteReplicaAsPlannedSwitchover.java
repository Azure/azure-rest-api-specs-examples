
import com.azure.resourcemanager.postgresqlflexibleserver.models.ReadReplicaPromoteMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Replica;
import com.azure.resourcemanager.postgresqlflexibleserver.models.ReplicationPromoteOption;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Server;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Servers Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * PromoteReplicaAsPlannedSwitchover.json
     */
    /**
     * Sample code: SwitchOver a replica server as planned, i.e. it will wait for replication to complete before
     * promoting replica as Primary and original primary as replica.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        switchOverAReplicaServerAsPlannedIEItWillWaitForReplicationToCompleteBeforePromotingReplicaAsPrimaryAndOriginalPrimaryAsReplica(
            com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Server resource = manager.servers()
            .getByResourceGroupWithResponse("testResourceGroup", "pgtestsvc4-replica", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withReplica(new Replica().withPromoteMode(ReadReplicaPromoteMode.SWITCHOVER)
            .withPromoteOption(ReplicationPromoteOption.PLANNED)).apply();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
