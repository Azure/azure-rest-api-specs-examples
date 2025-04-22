
import com.azure.resourcemanager.neonpostgres.models.Attributes;
import com.azure.resourcemanager.neonpostgres.models.NeonDatabaseProperties;
import java.util.Arrays;

/**
 * Samples for NeonDatabases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/NeonDatabases_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: NeonDatabases_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void
        neonDatabasesCreateOrUpdateMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.neonDatabases().define("entity-name")
            .withExistingBranche("rgneon", "test-org", "entity-name", "entity-name")
            .withProperties(new NeonDatabaseProperties().withEntityName("entity-name")
                .withAttributes(Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
                .withBranchId("orfdwdmzvfvlnrgussvcvoek").withOwnerName("odmbeg"))
            .create();
    }
}
