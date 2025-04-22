
import com.azure.resourcemanager.neonpostgres.models.Attributes;
import com.azure.resourcemanager.neonpostgres.models.NeonRole;
import com.azure.resourcemanager.neonpostgres.models.NeonRoleProperties;
import java.util.Arrays;

/**
 * Samples for NeonRoles Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/NeonRoles_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: NeonRoles_Update_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void neonRolesUpdateMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        NeonRole resource = manager.neonRoles().getWithResponse("rgneon", "test-org", "entity-name", "entity-name",
            "entity-name", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new NeonRoleProperties().withEntityName("entity-name")
                .withAttributes(Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
                .withBranchId("wxbojkmdgaggkfiwqfakdkbyztm")
                .withPermissions(Arrays.asList("myucqecpjriewzohxvadgkhiudnyx")).withIsSuperUser(true))
            .apply();
    }
}
