
/**
 * Samples for FluidRelayServers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_Get.
     * json
     */
    /**
     * Sample code: Get Fluid Relay server details.
     * 
     * @param manager Entry point to FluidRelayManager.
     */
    public static void getFluidRelayServerDetails(com.azure.resourcemanager.fluidrelay.FluidRelayManager manager) {
        manager.fluidRelayServers().getByResourceGroupWithResponse("myResourceGroup", "myFluidRelayServer",
            com.azure.core.util.Context.NONE);
    }
}
