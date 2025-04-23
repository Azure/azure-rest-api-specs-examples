
import com.azure.resourcemanager.neonpostgres.models.Attributes;
import com.azure.resourcemanager.neonpostgres.models.NeonRoleProperties;
import java.util.Arrays;

/**
 * Samples for NeonRoles CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/NeonRoles_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: NeonRoles_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void
        neonRolesCreateOrUpdateMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.neonRoles().define("entity-name")
            .withExistingBranche("rgneon", "test-org", "entity-name", "entity-name")
            .withProperties(new NeonRoleProperties().withEntityName("entity-name")
                .withAttributes(Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
                .withBranchId("wxbojkmdgaggkfiwqfakdkbyztm")
                .withPermissions(Arrays.asList("myucqecpjriewzohxvadgkhiudnyx")).withIsSuperUser(true))
            .create();
    }
}
