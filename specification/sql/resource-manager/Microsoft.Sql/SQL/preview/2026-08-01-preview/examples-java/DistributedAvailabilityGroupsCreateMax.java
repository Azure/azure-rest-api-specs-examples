
import com.azure.resourcemanager.sql.fluent.models.DistributedAvailabilityGroupInner;
import com.azure.resourcemanager.sql.models.DistributedAvailabilityGroupDatabase;
import com.azure.resourcemanager.sql.models.FailoverModeType;
import com.azure.resourcemanager.sql.models.LinkModeType;
import com.azure.resourcemanager.sql.models.LinkRole;
import com.azure.resourcemanager.sql.models.SeedingModeType;
import java.util.Arrays;

/**
 * Samples for DistributedAvailabilityGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/DistributedAvailabilityGroupsCreateMax.json
     */
    /**
     * Sample code: Create a distributed availability group with all properties.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createADistributedAvailabilityGroupWithAllProperties(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDistributedAvailabilityGroups().createOrUpdate("testrg", "testcl", "dag",
            new DistributedAvailabilityGroupInner().withPartnerAvailabilityGroupName("BoxLocalAg1")
                .withPartnerEndpoint("TCP://SERVER:7022").withInstanceLinkRole(LinkRole.PRIMARY)
                .withInstanceAvailabilityGroupName("testcl").withFailoverMode(FailoverModeType.NONE)
                .withSeedingMode(SeedingModeType.AUTOMATIC).withLinkMode(LinkModeType.SINGLE_DATABASE)
                .withDatabases(Arrays.asList(new DistributedAvailabilityGroupDatabase().withDatabaseName("testdb"))),
            com.azure.core.util.Context.NONE);
    }
}
