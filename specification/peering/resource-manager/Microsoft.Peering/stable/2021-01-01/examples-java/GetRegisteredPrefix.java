
/**
 * Samples for RegisteredPrefixes Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/GetRegisteredPrefix.json
     */
    /**
     * Sample code: Get a registered prefix associated with the peering.
     * 
     * @param manager Entry point to PeeringManager.
     */
    public static void
        getARegisteredPrefixAssociatedWithThePeering(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.registeredPrefixes().getWithResponse("rgName", "peeringName", "registeredPrefixName",
            com.azure.core.util.Context.NONE);
    }
}
