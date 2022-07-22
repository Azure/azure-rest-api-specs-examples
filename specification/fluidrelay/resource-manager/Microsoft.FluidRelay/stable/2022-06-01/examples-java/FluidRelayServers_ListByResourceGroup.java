import com.azure.core.util.Context;

/** Samples for FluidRelayServers ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_ListByResourceGroup.json
     */
    /**
     * Sample code: List all Fluid Relay servers in a resource group.
     *
     * @param manager Entry point to FluidRelayManager.
     */
    public static void listAllFluidRelayServersInAResourceGroup(
        com.azure.resourcemanager.fluidrelay.FluidRelayManager manager) {
        manager.fluidRelayServers().listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
