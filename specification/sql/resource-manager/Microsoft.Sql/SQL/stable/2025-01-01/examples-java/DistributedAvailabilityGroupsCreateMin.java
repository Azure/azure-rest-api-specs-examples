
import com.azure.resourcemanager.sql.fluent.models.DistributedAvailabilityGroupInner;
import com.azure.resourcemanager.sql.models.DistributedAvailabilityGroupDatabase;
import java.util.Arrays;

/**
 * Samples for DistributedAvailabilityGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DistributedAvailabilityGroupsCreateMin.json
     */
    /**
     * Sample code: Create a distributed availability group with minimal properties.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createADistributedAvailabilityGroupWithMinimalProperties(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDistributedAvailabilityGroups().createOrUpdate("testrg", "testcl", "dag",
            new DistributedAvailabilityGroupInner().withPartnerAvailabilityGroupName("BoxLocalAg1")
                .withPartnerEndpoint("TCP://SERVER:7022").withInstanceAvailabilityGroupName("testcl")
                .withDatabases(Arrays.asList(new DistributedAvailabilityGroupDatabase().withDatabaseName("testdb"))),
            com.azure.core.util.Context.NONE);
    }
}
