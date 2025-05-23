
import com.azure.resourcemanager.neonpostgres.models.Attributes;
import com.azure.resourcemanager.neonpostgres.models.EndpointProperties;
import com.azure.resourcemanager.neonpostgres.models.EndpointType;
import java.util.Arrays;

/**
 * Samples for Endpoints CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Endpoints_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: Endpoints_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void
        endpointsCreateOrUpdateMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.endpoints().define("entity-name")
            .withExistingBranche("rgneon", "test-org", "entity-name", "entity-name")
            .withProperties(new EndpointProperties().withEntityName("entity-name")
                .withAttributes(Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
                .withProjectId("rtvdeeflqzlrpfzhjqhcsfbldw").withBranchId("rzsyrhpfbydxtfkpaa")
                .withEndpointType(EndpointType.READ_ONLY))
            .create();
    }
}
