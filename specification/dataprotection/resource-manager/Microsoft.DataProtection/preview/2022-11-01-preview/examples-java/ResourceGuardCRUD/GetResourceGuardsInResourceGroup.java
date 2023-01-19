/** Samples for ResourceGuards ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/ResourceGuardCRUD/GetResourceGuardsInResourceGroup.json
     */
    /**
     * Sample code: Get ResourceGuards in ResourceGroup.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getResourceGuardsInResourceGroup(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.resourceGuards().listByResourceGroup("SampleResourceGroup", com.azure.core.util.Context.NONE);
    }
}
