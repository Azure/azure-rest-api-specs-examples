
/**
 * Samples for Gates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/Gates_Get.json
     */
    /**
     * Sample code: Gets a Gate resource.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void
        getsAGateResource(com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.gates().getWithResponse("rg1", "fleet1", "12345678-910a-bcde-f000-000000000000",
            com.azure.core.util.Context.NONE);
    }
}
