
/**
 * Samples for Organizations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-01-preview/Organizations_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Get.
     * 
     * @param manager Entry point to NeonPostgresManager.
     */
    public static void organizationsGet(com.azure.resourcemanager.neonpostgres.NeonPostgresManager manager) {
        manager.organizations().getByResourceGroupWithResponse("rgneon", "5", com.azure.core.util.Context.NONE);
    }
}
