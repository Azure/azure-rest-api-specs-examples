
import com.azure.resourcemanager.neonpostgres.models.Attributes;
import com.azure.resourcemanager.neonpostgres.models.Branch;
import com.azure.resourcemanager.neonpostgres.models.BranchProperties;
import com.azure.resourcemanager.neonpostgres.models.EndpointProperties;
import com.azure.resourcemanager.neonpostgres.models.EndpointType;
import com.azure.resourcemanager.neonpostgres.models.NeonDatabaseProperties;
import com.azure.resourcemanager.neonpostgres.models.NeonRoleProperties;
import java.util.Arrays;

/**
 * Samples for Branches Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Branches_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: Branches_Update_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void branchesUpdateMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        Branch resource = manager.branches()
            .getWithResponse("rgneon", "test-org", "entity-name", "entity-name", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new BranchProperties().withEntityName("entity-name")
            .withAttributes(Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
            .withProjectId("oik").withParentId("entity-id").withRoleName("qrrairsupyosxnqotdwhbpc")
            .withDatabaseName("duhxebzhd")
            .withRoles(Arrays.asList(new NeonRoleProperties().withEntityName("entity-name")
                .withAttributes(Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
                .withBranchId("wxbojkmdgaggkfiwqfakdkbyztm")
                .withPermissions(Arrays.asList("myucqecpjriewzohxvadgkhiudnyx")).withIsSuperUser(true)))
            .withDatabases(Arrays.asList(new NeonDatabaseProperties().withEntityName("entity-name")
                .withAttributes(Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
                .withBranchId("orfdwdmzvfvlnrgussvcvoek").withOwnerName("odmbeg")))
            .withEndpoints(Arrays.asList(new EndpointProperties().withEntityName("entity-name")
                .withAttributes(Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
                .withProjectId("rtvdeeflqzlrpfzhjqhcsfbldw").withBranchId("rzsyrhpfbydxtfkpaa")
                .withEndpointType(EndpointType.READ_ONLY))))
            .apply();
    }
}
