
import com.azure.resourcemanager.neonpostgres.models.Attributes;
import com.azure.resourcemanager.neonpostgres.models.BranchProperties;
import com.azure.resourcemanager.neonpostgres.models.DefaultEndpointSettings;
import com.azure.resourcemanager.neonpostgres.models.EndpointProperties;
import com.azure.resourcemanager.neonpostgres.models.EndpointType;
import com.azure.resourcemanager.neonpostgres.models.NeonDatabaseProperties;
import com.azure.resourcemanager.neonpostgres.models.NeonRoleProperties;
import com.azure.resourcemanager.neonpostgres.models.ProjectProperties;
import java.util.Arrays;

/**
 * Samples for Projects CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Projects_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: Projects_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void
        projectsCreateOrUpdateMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.projects().define("entity-name").withExistingOrganization("rgneon", "test-org")
            .withProperties(new ProjectProperties().withEntityName("entity-name")
                .withAttributes(Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
                .withRegionId("tlcltldfrnxh").withStorage(7L).withPgVersion(10).withHistoryRetention(7)
                .withDefaultEndpointSettings(
                    new DefaultEndpointSettings().withAutoscalingLimitMinCu(26.0).withAutoscalingLimitMaxCu(20.0))
                .withBranch(new BranchProperties().withEntityName("entity-name")
                    .withAttributes(
                        Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
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
