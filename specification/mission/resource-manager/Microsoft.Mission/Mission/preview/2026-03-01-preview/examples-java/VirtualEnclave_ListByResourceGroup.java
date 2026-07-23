
/**
 * Samples for VirtualEnclave ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/VirtualEnclave_ListByResourceGroup.json
     */
    /**
     * Sample code: VirtualEnclave_ListByResourceGroup.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        virtualEnclaveListByResourceGroup(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.virtualEnclaves().listByResourceGroup("rgopenapi", com.azure.core.util.Context.NONE);
    }
}
