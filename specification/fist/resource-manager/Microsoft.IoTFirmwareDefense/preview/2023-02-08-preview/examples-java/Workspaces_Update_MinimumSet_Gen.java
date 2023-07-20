import com.azure.resourcemanager.iotfirmwaredefense.models.Workspace;

/** Samples for Workspaces Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Workspaces_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_Update_MinimumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void workspacesUpdateMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        Workspace resource =
            manager
                .workspaces()
                .getByResourceGroupWithResponse("rgworkspaces", "E___-3", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
