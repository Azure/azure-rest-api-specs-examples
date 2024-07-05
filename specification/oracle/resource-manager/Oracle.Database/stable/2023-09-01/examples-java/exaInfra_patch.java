
import com.azure.resourcemanager.oracledatabase.models.CloudExadataInfrastructure;

/**
 * Samples for CloudExadataInfrastructures Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/exaInfra_patch.json
     */
    /**
     * Sample code: Patch Exadata Infrastructure.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        patchExadataInfrastructure(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        CloudExadataInfrastructure resource = manager.cloudExadataInfrastructures()
            .getByResourceGroupWithResponse("rg000", "infra1", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
