
/**
 * Samples for Workspaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/
     * Workspaces_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void
        workspacesDeleteMinimumSetGen(com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.workspaces().deleteByResourceGroupWithResponse("rgworkspaces", "E___-3",
            com.azure.core.util.Context.NONE);
    }
}
