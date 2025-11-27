
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
     * ServersPromoteReplicaAsPlannedSwitchover.json
     */
    /**
     * Sample code: Switch over a read replica to primary server with planned data synchronization. Meaning that it
     * waits for data in the read replica to be fully synchronized with its source server before it initiates the
     * switching of roles between the read replica and the primary server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        switchOverAReadReplicaToPrimaryServerWithPlannedDataSynchronizationMeaningThatItWaitsForDataInTheReadReplicaToBeFullySynchronizedWithItsSourceServerBeforeItInitiatesTheSwitchingOfRolesBetweenTheReadReplicaAndThePrimaryServer(
            com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Server resource = manager.servers()
            .getByResourceGroupWithResponse("exampleresourcegroup", "exampleserver", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withReplica(new Replica().withPromoteMode(ReadReplicaPromoteMode.SWITCHOVER)
            .withPromoteOption(ReadReplicaPromoteOption.PLANNED)).apply();
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
