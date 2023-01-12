/** Samples for Registrations GetActivationKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/Registration/Post.json
     */
    /**
     * Sample code: Returns Azure Stack Activation Key.
     *
     * @param manager Entry point to AzureStackManager.
     */
    public static void returnsAzureStackActivationKey(com.azure.resourcemanager.azurestack.AzureStackManager manager) {
        manager
            .registrations()
            .getActivationKeyWithResponse("azurestack", "testregistration", com.azure.core.util.Context.NONE);
    }
}
