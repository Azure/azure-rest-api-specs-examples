
import com.azure.resourcemanager.neonpostgres.models.Attributes;
import com.azure.resourcemanager.neonpostgres.models.BranchProperties;
import com.azure.resourcemanager.neonpostgres.models.EndpointProperties;
import com.azure.resourcemanager.neonpostgres.models.EndpointType;
import com.azure.resourcemanager.neonpostgres.models.NeonDatabaseProperties;
import com.azure.resourcemanager.neonpostgres.models.NeonRoleProperties;
import java.util.Arrays;

/**
 * Samples for Branches CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Branches_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: Branches_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void
        branchesCreateOrUpdateMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.branches().define("test-entity").withExistingProject("rgneon", "test-org", "test-entity")
            .withProperties(new BranchProperties().withEntityName("entity-name")
                .withAttributes(Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
                .withProjectId("oik").withParentId("entity-id").withRoleName("qrrairsupyosxnqotdwhbpc")
                .withDatabaseName("duhxebzhd")
                .withRoles(Arrays.asList(new NeonRoleProperties().withEntityName("entity-name")
                    .withAttributes(
                        Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
                    .withBranchId("wxbojkmdgaggkfiwqfakdkbyztm")
                    .withPermissions(Arrays.asList("myucqecpjriewzohxvadgkhiudnyx")).withIsSuperUser(true)))
                .withDatabases(Arrays.asList(new NeonDatabaseProperties().withEntityName("entity-name")
                    .withAttributes(
                        Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
                    .withBranchId("orfdwdmzvfvlnrgussvcvoek").withOwnerName("odmbeg")))
                .withEndpoints(Arrays.asList(new EndpointProperties().withEntityName("entity-name")
                    .withAttributes(
                        Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
                    .withProjectId("rtvdeeflqzlrpfzhjqhcsfbldw").withBranchId("rzsyrhpfbydxtfkpaa")
                    .withEndpointType(EndpointType.READ_ONLY))))
            .create();
    }
}
