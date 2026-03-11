
/**
 * Samples for HardwareSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/HardwareSettings_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: HardwareSettings_Get_MaximumSet.
     * 
     * @param manager Entry point to DisconnectedOperationsManager.
     */
    public static void hardwareSettingsGetMaximumSet(
        com.azure.resourcemanager.disconnectedoperations.DisconnectedOperationsManager manager) {
        manager.hardwareSettings().getWithResponse("rgdisconnectedOperations", "demo-resource", "default",
            com.azure.core.util.Context.NONE);
    }
}
