/** Samples for Registrations EnableRemoteManagement. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/RemoteManagement/Post.json
     */
    /**
     * Sample code: Returns empty response for successful action..
     *
     * @param manager Entry point to AzureStackManager.
     */
    public static void returnsEmptyResponseForSuccessfulAction(
        com.azure.resourcemanager.azurestack.AzureStackManager manager) {
        manager
            .registrations()
            .enableRemoteManagementWithResponse("azurestack", "testregistration", com.azure.core.util.Context.NONE);
    }
}
