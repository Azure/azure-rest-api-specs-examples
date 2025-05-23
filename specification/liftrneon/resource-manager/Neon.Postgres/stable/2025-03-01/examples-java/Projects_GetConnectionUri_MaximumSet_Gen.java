
import com.azure.resourcemanager.neonpostgres.fluent.models.ConnectionUriPropertiesInner;

/**
 * Samples for Projects GetConnectionUri.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Projects_GetConnectionUri_MaximumSet_Gen.json
     */
    /**
     * Sample code: Projects_GetConnectionUri_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void
        projectsGetConnectionUriMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.projects().getConnectionUriWithResponse("rgneon", "test-org", "entity-name",
            new ConnectionUriPropertiesInner().withProjectId("riuifmoqtorrcffgksvfcobia").withBranchId("iimmlbqv")
                .withDatabaseName("xc").withRoleName("xhmcvsgtp").withEndpointId("jcpdvsyjcn").withIsPooled(true),
            com.azure.core.util.Context.NONE);
    }
}
