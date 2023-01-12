/** Samples for Registrations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/Registration/Delete.json
     */
    /**
     * Sample code: Delete the requested Azure Stack registration.
     *
     * @param manager Entry point to AzureStackManager.
     */
    public static void deleteTheRequestedAzureStackRegistration(
        com.azure.resourcemanager.azurestack.AzureStackManager manager) {
        manager
            .registrations()
            .deleteByResourceGroupWithResponse("azurestack", "testregistration", com.azure.core.util.Context.NONE);
    }
}
