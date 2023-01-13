/** Samples for Registrations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/Registration/ListBySubscription.json
     */
    /**
     * Sample code: Returns a list of all registrations under current subscription.
     *
     * @param manager Entry point to AzureStackManager.
     */
    public static void returnsAListOfAllRegistrationsUnderCurrentSubscription(
        com.azure.resourcemanager.azurestack.AzureStackManager manager) {
        manager.registrations().list(com.azure.core.util.Context.NONE);
    }
}
