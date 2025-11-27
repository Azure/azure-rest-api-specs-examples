
import com.azure.resourcemanager.postgresqlflexibleserver.models.ReadReplicaPromoteMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.ReadReplicaPromoteOption;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Replica;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Server;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Servers Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * ServersPromoteReplicaAsForcedStandaloneServer.json
     */
    /**
     * Sample code: Promote a read replica to a standalone server with forced data synchronization. Meaning that it
     * doesn't wait for data in the read replica to be synchronized with its source server before it initiates the
     * promotion to a standalone server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        promoteAReadReplicaToAStandaloneServerWithForcedDataSynchronizationMeaningThatItDoesnTWaitForDataInTheReadReplicaToBeSynchronizedWithItsSourceServerBeforeItInitiatesThePromotionToAStandaloneServer(
            com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Server resource = manager.servers()
            .getByResourceGroupWithResponse("exampleresourcegroup", "exampleserver", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withReplica(new Replica().withPromoteMode(ReadReplicaPromoteMode.STANDALONE)
            .withPromoteOption(ReadReplicaPromoteOption.FORCED)).apply();
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
