/** Samples for ApplicationDefinitions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/listApplicationDefinitionsBySubscription.json
     */
    /**
     * Sample code: Lists all the application definitions within a subscription.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void listsAllTheApplicationDefinitionsWithinASubscription(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.applicationDefinitions().list(com.azure.core.util.Context.NONE);
    }
}
