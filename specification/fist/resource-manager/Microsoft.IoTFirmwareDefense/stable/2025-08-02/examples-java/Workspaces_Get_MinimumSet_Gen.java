
/**
 * Samples for Workspaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-02/Workspaces_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_Get_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void
        workspacesGetMinimumSetGen(com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.workspaces().getByResourceGroupWithResponse("rgworkspaces", "E_US", com.azure.core.util.Context.NONE);
    }
}
