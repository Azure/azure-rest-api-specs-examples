
/**
 * Samples for ResourceGuards ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * ResourceGuardCRUD/GetResourceGuardsInResourceGroup.json
     */
    /**
     * Sample code: Get ResourceGuards in ResourceGroup.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        getResourceGuardsInResourceGroup(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.resourceGuards().listByResourceGroup("SampleResourceGroup", com.azure.core.util.Context.NONE);
    }
}
