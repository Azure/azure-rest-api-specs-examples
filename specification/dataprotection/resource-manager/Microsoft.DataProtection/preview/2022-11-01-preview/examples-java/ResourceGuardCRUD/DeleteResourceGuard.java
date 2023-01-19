/** Samples for ResourceGuards Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/ResourceGuardCRUD/DeleteResourceGuard.json
     */
    /**
     * Sample code: Delete ResourceGuard.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void deleteResourceGuard(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .resourceGuards()
            .deleteByResourceGroupWithResponse(
                "SampleResourceGroup", "swaggerExample", com.azure.core.util.Context.NONE);
    }
}
