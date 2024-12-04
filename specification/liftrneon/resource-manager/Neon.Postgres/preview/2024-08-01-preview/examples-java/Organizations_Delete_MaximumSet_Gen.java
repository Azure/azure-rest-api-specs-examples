
/**
 * Samples for Organizations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-01-preview/Organizations_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Delete.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void organizationsDelete(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.organizations().delete("rgneon", "2_3", com.azure.core.util.Context.NONE);
    }
}
