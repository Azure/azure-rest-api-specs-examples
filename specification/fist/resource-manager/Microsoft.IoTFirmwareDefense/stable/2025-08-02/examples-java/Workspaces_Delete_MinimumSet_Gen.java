
/**
 * Samples for Workspaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-02/Workspaces_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void
        workspacesDeleteMinimumSetGen(com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.workspaces().delete("rgworkspaces", "default", com.azure.core.util.Context.NONE);
    }
}
