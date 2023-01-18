import com.azure.resourcemanager.powerbidedicated.models.CheckCapacityNameAvailabilityParameters;

/** Samples for Capacities CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/checkNameAvailability.json
     */
    /**
     * Sample code: Check name availability of a capacity.
     *
     * @param manager Entry point to PowerBIDedicatedManager.
     */
    public static void checkNameAvailabilityOfACapacity(
        com.azure.resourcemanager.powerbidedicated.PowerBIDedicatedManager manager) {
        manager
            .capacities()
            .checkNameAvailabilityWithResponse(
                "West US",
                new CheckCapacityNameAvailabilityParameters()
                    .withName("azsdktest")
                    .withType("Microsoft.PowerBIDedicated/capacities"),
                com.azure.core.util.Context.NONE);
    }
}
