
/**
 * Samples for ResourceGuards GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ResourceGuardCRUD/GetResourceGuard.json
     */
    /**
     * Sample code: Get ResourceGuard.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getResourceGuard(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.resourceGuards().getByResourceGroupWithResponse("SampleResourceGroup", "swaggerExample",
            com.azure.core.util.Context.NONE);
    }
}
