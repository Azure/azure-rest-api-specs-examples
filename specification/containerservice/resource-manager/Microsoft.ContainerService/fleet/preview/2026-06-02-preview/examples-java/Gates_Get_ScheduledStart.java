
/**
 * Samples for Gates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/Gates_Get_ScheduledStart.json
     */
    /**
     * Sample code: Gets a ScheduledStart Gate resource.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void getsAScheduledStartGateResource(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.gates().getWithResponse("rg1", "fleet1", "12345678-910a-bcde-f000-000000000001",
            com.azure.core.util.Context.NONE);
    }
}
