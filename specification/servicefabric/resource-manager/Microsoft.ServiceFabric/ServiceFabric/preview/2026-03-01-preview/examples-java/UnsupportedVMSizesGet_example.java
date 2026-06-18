
/**
 * Samples for UnsupportedVmSizes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/UnsupportedVMSizesGet_example.json
     */
    /**
     * Sample code: Get unsupported vm sizes.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void getUnsupportedVmSizes(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.unsupportedVmSizes().getWithResponse("eastus", "Standard_B1ls1", com.azure.core.util.Context.NONE);
    }
}
