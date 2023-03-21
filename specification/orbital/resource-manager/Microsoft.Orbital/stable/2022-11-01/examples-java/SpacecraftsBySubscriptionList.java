/** Samples for Spacecrafts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/SpacecraftsBySubscriptionList.json
     */
    /**
     * Sample code: List of Spacecraft by Subscription.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void listOfSpacecraftBySubscription(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager.spacecrafts().list("opaqueString", com.azure.core.util.Context.NONE);
    }
}
