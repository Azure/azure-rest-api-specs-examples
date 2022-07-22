import com.azure.core.util.Context;

/** Samples for FluidRelayServers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_ListBySubscription.json
     */
    /**
     * Sample code: List all Fluid Relay servers in a subscription.
     *
     * @param manager Entry point to FluidRelayManager.
     */
    public static void listAllFluidRelayServersInASubscription(
        com.azure.resourcemanager.fluidrelay.FluidRelayManager manager) {
        manager.fluidRelayServers().list(Context.NONE);
    }
}
