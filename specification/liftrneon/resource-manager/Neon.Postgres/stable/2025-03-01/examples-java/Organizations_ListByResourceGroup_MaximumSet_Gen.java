
/**
 * Samples for Organizations ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Organizations_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void
        organizationsListByResourceGroupMaximumSet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.organizations().listByResourceGroup("rgneon", com.azure.core.util.Context.NONE);
    }
}
