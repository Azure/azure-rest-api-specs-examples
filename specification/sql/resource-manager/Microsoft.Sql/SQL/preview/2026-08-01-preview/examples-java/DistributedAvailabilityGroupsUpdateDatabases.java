
import com.azure.resourcemanager.sql.fluent.models.DistributedAvailabilityGroupInner;
import com.azure.resourcemanager.sql.models.DistributedAvailabilityGroupDatabase;
import java.util.Arrays;

/**
 * Samples for DistributedAvailabilityGroups Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/DistributedAvailabilityGroupsUpdateDatabases.json
     */
    /**
     * Sample code: Update the databases of a distributed availability group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        updateTheDatabasesOfADistributedAvailabilityGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDistributedAvailabilityGroups().update("testrg", "testcl", "dag",
            new DistributedAvailabilityGroupInner()
                .withDatabases(Arrays.asList(new DistributedAvailabilityGroupDatabase().withDatabaseName("testdb1"),
                    new DistributedAvailabilityGroupDatabase().withDatabaseName("testdb2"),
                    new DistributedAvailabilityGroupDatabase().withDatabaseName("testdb3"))),
            com.azure.core.util.Context.NONE);
    }
}
