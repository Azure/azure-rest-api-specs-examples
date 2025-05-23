
import com.azure.resourcemanager.neonpostgres.models.Attributes;
import com.azure.resourcemanager.neonpostgres.models.Compute;
import com.azure.resourcemanager.neonpostgres.models.ComputeProperties;
import java.util.Arrays;

/**
 * Samples for Computes Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Computes_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: Computes_Update_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void computesUpdateMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        Compute resource = manager.computes().getWithResponse("rgneon", "test-org", "entity-name", "entity-name",
            "entity-name", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new ComputeProperties().withEntityName("entity-name")
                .withAttributes(Arrays.asList(new Attributes().withName("trhvzyvaqy").withValue("evpkgsskyavybxwwssm")))
                .withRegion("mcfyojzptdliawyuxyxzqxif").withCpuCores(29).withMemory(2).withStatus("upwdpznysuwt"))
            .apply();
    }
}
