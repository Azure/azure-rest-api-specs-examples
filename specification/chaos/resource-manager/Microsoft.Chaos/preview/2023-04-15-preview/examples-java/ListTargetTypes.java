/** Samples for TargetTypes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/ListTargetTypes.json
     */
    /**
     * Sample code: List all Target Types for westus2 location.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void listAllTargetTypesForWestus2Location(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.targetTypes().list("westus2", null, com.azure.core.util.Context.NONE);
    }
}
