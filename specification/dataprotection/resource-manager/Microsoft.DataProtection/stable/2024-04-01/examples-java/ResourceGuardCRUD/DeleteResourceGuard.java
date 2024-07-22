
/**
 * Samples for ResourceGuards Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * ResourceGuardCRUD/DeleteResourceGuard.json
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
