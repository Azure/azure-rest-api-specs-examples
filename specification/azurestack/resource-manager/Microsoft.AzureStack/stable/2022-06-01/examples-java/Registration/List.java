/** Samples for Registrations ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/Registration/List.json
     */
    /**
     * Sample code: Returns a list of all registrations.
     *
     * @param manager Entry point to AzureStackManager.
     */
    public static void returnsAListOfAllRegistrations(com.azure.resourcemanager.azurestack.AzureStackManager manager) {
        manager.registrations().listByResourceGroup("azurestack", com.azure.core.util.Context.NONE);
    }
}
