
/**
 * Samples for ResourceGuards Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ResourceGuardCRUD/DeleteResourceGuard.json
     */
    /**
     * Sample code: Delete ResourceGuard.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void deleteResourceGuard(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.resourceGuards().deleteByResourceGroupWithResponse("SampleResourceGroup", "swaggerExample",
            com.azure.core.util.Context.NONE);
    }
}
