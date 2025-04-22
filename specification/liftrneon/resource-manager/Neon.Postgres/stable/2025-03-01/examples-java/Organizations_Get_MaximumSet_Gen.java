
/**
 * Samples for Organizations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Organizations_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Get_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void organizationsGetMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.organizations().getByResourceGroupWithResponse("rgneon", "test-org", com.azure.core.util.Context.NONE);
    }
}
