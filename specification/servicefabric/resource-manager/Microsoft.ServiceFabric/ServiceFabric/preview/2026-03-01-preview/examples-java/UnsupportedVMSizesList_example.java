
/**
 * Samples for UnsupportedVmSizes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/UnsupportedVMSizesList_example.json
     */
    /**
     * Sample code: List unsupported vm sizes.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void listUnsupportedVmSizes(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.unsupportedVmSizes().list("eastus", com.azure.core.util.Context.NONE);
    }
}
