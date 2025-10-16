
/**
 * Samples for ResourceGuards ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ResourceGuardCRUD/GetResourceGuardsInResourceGroup.json
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
