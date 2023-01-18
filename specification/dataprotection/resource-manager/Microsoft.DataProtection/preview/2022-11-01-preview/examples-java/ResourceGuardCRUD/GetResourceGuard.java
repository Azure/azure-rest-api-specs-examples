/** Samples for ResourceGuards GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/ResourceGuardCRUD/GetResourceGuard.json
     */
    /**
     * Sample code: Get ResourceGuard.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getResourceGuard(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .resourceGuards()
            .getByResourceGroupWithResponse("SampleResourceGroup", "swaggerExample", com.azure.core.util.Context.NONE);
    }
}
