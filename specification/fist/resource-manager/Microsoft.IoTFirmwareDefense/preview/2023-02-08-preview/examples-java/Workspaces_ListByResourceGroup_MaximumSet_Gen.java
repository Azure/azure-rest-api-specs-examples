/** Samples for Workspaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Workspaces_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_ListByResourceGroup_MaximumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void workspacesListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.workspaces().listByResourceGroup("rgworkspaces", com.azure.core.util.Context.NONE);
    }
}
