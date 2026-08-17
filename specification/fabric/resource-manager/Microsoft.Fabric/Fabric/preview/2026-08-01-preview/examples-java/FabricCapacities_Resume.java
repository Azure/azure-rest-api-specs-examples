
/**
 * Samples for FabricCapacities Resume.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/FabricCapacities_Resume.json
     */
    /**
     * Sample code: Resume capacity.
     * 
     * @param manager Entry point to FabricManager.
     */
    public static void resumeCapacity(com.azure.resourcemanager.fabric.FabricManager manager) {
        manager.fabricCapacities().resume("TestRG", "azsdktest", com.azure.core.util.Context.NONE);
    }
}
