
/**
 * Samples for Workspaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/
     * Workspaces_ListByResourceGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_ListByResourceGroup_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void workspacesListByResourceGroupMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.workspaces().listByResourceGroup("rgworkspaces", com.azure.core.util.Context.NONE);
    }
}
