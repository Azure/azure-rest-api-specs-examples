import com.azure.core.util.Context;
import com.azure.resourcemanager.signalr.models.NameAvailabilityParameters;

/** Samples for SignalR CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_CheckNameAvailability.json
     */
    /**
     * Sample code: SignalR_CheckNameAvailability.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRCheckNameAvailability(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRs()
            .checkNameAvailabilityWithResponse(
                "eastus",
                new NameAvailabilityParameters()
                    .withType("Microsoft.SignalRService/SignalR")
                    .withName("mySignalRService"),
                Context.NONE);
    }
}
