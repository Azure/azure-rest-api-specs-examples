
/**
 * Samples for HardwareSettings ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/HardwareSettings_ListByParent_MaximumSet_Gen.json
     */
    /**
     * Sample code: HardwareSettings_ListByParent_MaximumSet.
     * 
     * @param manager Entry point to DisconnectedOperationsManager.
     */
    public static void hardwareSettingsListByParentMaximumSet(
        com.azure.resourcemanager.disconnectedoperations.DisconnectedOperationsManager manager) {
        manager.hardwareSettings().listByParent("rgdisconnectedOperations", "demo-resource",
            com.azure.core.util.Context.NONE);
    }
}
