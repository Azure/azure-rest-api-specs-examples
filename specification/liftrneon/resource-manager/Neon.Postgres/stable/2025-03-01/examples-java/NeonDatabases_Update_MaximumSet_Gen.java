
import com.azure.resourcemanager.neonpostgres.models.Attributes;
import com.azure.resourcemanager.neonpostgres.models.NeonDatabase;
import com.azure.resourcemanager.neonpostgres.models.NeonDatabaseProperties;
import java.util.Arrays;

/**
 * Samples for NeonDatabases Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/NeonDatabases_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: NeonDatabases_Update_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void
        neonDatabasesUpdateMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        NeonDatabase resource = manager.neonDatabases().getWithResponse("rgneon", "test-org", "entity-name",
            "entity-name", "entity-name", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new NeonDatabaseProperties().withEntityName("entity-name")
                .withAttributes(Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
                .withBranchId("orfdwdmzvfvlnrgussvcvoek").withOwnerName("odmbeg"))
            .apply();
    }
}
