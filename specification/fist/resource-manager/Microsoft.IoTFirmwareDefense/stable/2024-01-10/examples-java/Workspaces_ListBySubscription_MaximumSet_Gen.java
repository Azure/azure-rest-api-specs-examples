
/**
 * Samples for Workspaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/
     * Workspaces_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void workspacesListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.workspaces().list(com.azure.core.util.Context.NONE);
    }
}
