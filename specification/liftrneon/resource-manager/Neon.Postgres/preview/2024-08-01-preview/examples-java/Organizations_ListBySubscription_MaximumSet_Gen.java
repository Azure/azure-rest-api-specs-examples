
/**
 * Samples for Organizations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-01-preview/Organizations_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_ListBySubscription.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void
        organizationsListBySubscription(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.organizations().list(com.azure.core.util.Context.NONE);
    }
}
