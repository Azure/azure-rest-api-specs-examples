
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
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-12-01-preview/examples/
     * PromoteReplicaAsForcedStandaloneServer.json
     */
    /**
     * Sample code: Promote a replica server as a Standalone server as forced, i.e. it will promote a replica server
     * immediately without waiting for primary and replica to be in sync.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        promoteAReplicaServerAsAStandaloneServerAsForcedIEItWillPromoteAReplicaServerImmediatelyWithoutWaitingForPrimaryAndReplicaToBeInSync(
            com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Server resource = manager.servers()
            .getByResourceGroupWithResponse("testResourceGroup", "pgtestsvc4-replica", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withReplica(new Replica().withPromoteMode(ReadReplicaPromoteMode.STANDALONE)
            .withPromoteOption(ReplicationPromoteOption.FORCED)).apply();
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
