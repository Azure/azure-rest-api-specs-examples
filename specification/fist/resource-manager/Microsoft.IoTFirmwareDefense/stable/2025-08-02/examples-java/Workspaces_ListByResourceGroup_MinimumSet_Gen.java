
/**
 * Samples for Workspaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-02/Workspaces_ListByResourceGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_ListByResourceGroup_MinimumSet.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void workspacesListByResourceGroupMinimumSet(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.workspaces().listByResourceGroup("rgiotfirmwaredefense", com.azure.core.util.Context.NONE);
    }
}
