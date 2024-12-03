
/**
 * Samples for Organizations ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-01-preview/Organizations_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_ListByResourceGroup.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void
        organizationsListByResourceGroup(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.organizations().listByResourceGroup("rgneon", com.azure.core.util.Context.NONE);
    }
}
