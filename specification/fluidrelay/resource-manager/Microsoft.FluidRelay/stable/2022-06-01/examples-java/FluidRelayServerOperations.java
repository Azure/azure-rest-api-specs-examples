import com.azure.core.util.Context;

/** Samples for FluidRelayOperations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServerOperations.json
     */
    /**
     * Sample code: List Fluid Relay server operations.
     *
     * @param manager Entry point to FluidRelayManager.
     */
    public static void listFluidRelayServerOperations(com.azure.resourcemanager.fluidrelay.FluidRelayManager manager) {
        manager.fluidRelayOperations().list(Context.NONE);
    }
}
